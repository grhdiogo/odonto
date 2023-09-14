package criteria

import (
	"bytes"
	"fmt"
	"strings"
)

type Criteria interface {
	And(k string, op OperatorType, v any) *criteriaBuilder
	Or(k string, op OperatorType, v any) *criteriaBuilder
	AndBlock(cns ...Condition) *criteriaBuilder
	OrBlock(cns ...Condition) *criteriaBuilder
	Build(startsIn int, withLimit bool) string
	SetPage(page int) *criteriaBuilder
	SetLimit(limit int) *criteriaBuilder
	AddGroupBy(field string) *criteriaBuilder
	SetSort(sort string) *criteriaBuilder
	Values(n ...any) []any
	BuildWithCount(tmpl string, startsIn int) string
	Clone() Criteria
}

type OperatorType string

const OPT_IS OperatorType = "is"
const OPT_EQUAL OperatorType = "="
const OPT_BIGGER OperatorType = ">"
const OPT_LOWER OperatorType = "<"
const OPT_ILIKE OperatorType = "ilike"
const OPT_LIKE OperatorType = "like"
const OPT_ISNULL OperatorType = "is null"

type ConditionType string

const (
	CT_OR  ConditionType = "or"
	CT_AND ConditionType = "and"
)

// ===================================================
// Where Builder
// ===================================================

type Condition interface {
	Exp(seq int) string
	Type() ConditionType
	Count() int
	Values() []any
}

// ===================================================
// Condition Impl
// ===================================================

// add value based on operator
func valueByOperator(value any, t OperatorType) any {
	switch t {
	case OPT_LIKE:
		return fmt.Sprintf("%%%%%s%%%%", value)
	case OPT_ILIKE:
		return fmt.Sprintf("%%%%%s%%%%", value)
	default:
		return value
	}
}

type conditionImpl struct {
	value any
	sql   string
	ct    ConditionType
}

func (cb conditionImpl) Exp(seq int) string {
	return fmt.Sprintf("%s $%d ", cb.sql, seq)
}

func (cb conditionImpl) Type() ConditionType {
	return cb.ct
}

func (cb conditionImpl) Count() int {
	return 1
}

func (cb conditionImpl) Values() []any {
	return []any{cb.value}
}

func NewAndCondition(k string, op OperatorType, v any) Condition {
	return conditionImpl{
		sql:   fmt.Sprintf("AND %s %s ", k, op),
		value: valueByOperator(v, op),
		ct:    CT_AND,
	}
}

func NewOrCondition(k string, op OperatorType, v any) Condition {
	return conditionImpl{
		sql:   fmt.Sprintf("OR %s %s ", k, op),
		value: valueByOperator(v, op),
		ct:    CT_OR,
	}
}

// ===================================================
// Block
// ===================================================

type othersConditionImpl struct {
	values []any
	sql    string
	ct     ConditionType
}

func (cb othersConditionImpl) buildBetween(seq int) string {
	return fmt.Sprintf("%s $%d AND $%d ", cb.sql, seq, seq+1)
}

func (cb othersConditionImpl) Exp(seq int) string {
	// only between for now
	return cb.buildBetween(seq)
}

func (cb othersConditionImpl) Type() ConditionType {
	return cb.ct
}

func (cb othersConditionImpl) Count() int {
	return len(cb.values)
}

func (cb othersConditionImpl) Values() []any {
	return cb.values
}

// range value between values v1 and v2
func NewBetweenCondition(ct ConditionType, k string, v1, v2 any) Condition {
	// and registration BETWEEN '08/11/2022' AND '09/11/2022';
	// ct k BEWEEN v1 AND v2
	return othersConditionImpl{
		sql:    fmt.Sprintf("%s %s BETWEEN", ct, k),
		values: []any{v1, v2},
		ct:     CT_AND,
	}
}

// ===================================================
// Block
// ===================================================

type blockCondition struct {
	ct    ConditionType
	conds []Condition
}

func (cb blockCondition) Count() int {
	c := 0
	for _, v := range cb.conds {
		c += v.Count()
	}
	return c
}

func (cb blockCondition) Exp(seq int) string {
	sql := bytes.Buffer{}
	internalSequence := seq
	// get type
	if cb.ct == CT_AND {
		sql.WriteString("AND ( ")
	} else {
		sql.WriteString("OR ( ")
	}
	//
	for i, v := range cb.conds {
		if i == 0 {
			// verify type
			t := v.Type()
			if t == CT_AND {
				sql.WriteString("1=1 ")
			} else {
				sql.WriteString("1!=1 ")
			}
		}
		sql.WriteString(v.Exp(internalSequence))
		internalSequence += v.Count()
	}
	sql.WriteString(") ")

	return sql.String()
}

func (cb blockCondition) Type() ConditionType {
	return cb.ct
}

func (cb blockCondition) Values() []any {
	result := make([]any, 0)
	for _, v := range cb.conds {
		result = append(result, v.Values()...)
	}
	return result
}

func NewAndBlockCondition(cns ...Condition) Condition {
	return blockCondition{
		conds: cns,
		ct:    CT_AND,
	}
}

func NewOrBlockCondition(cns ...Condition) Condition {
	return blockCondition{
		conds: cns,
		ct:    CT_OR,
	}
}

// ===================================================
// Builder
// ===================================================

type criteriaBuilder struct {
	limit     int
	pageIndex int
	sort      string
	groupBy   []string
	//
	filterValues []any
	conditions   []Condition
}

//
func (wb *criteriaBuilder) And(k string, op OperatorType, v any) *criteriaBuilder {
	// return wb.add(CT_AND, k, op, v)
	wb.filterValues = append(wb.filterValues, v)
	//
	wb.conditions = append(wb.conditions, NewAndCondition(k, op, v))
	return wb
}

func (wb *criteriaBuilder) Or(k string, op OperatorType, v any) *criteriaBuilder {
	wb.filterValues = append(wb.filterValues, v)
	//
	wb.conditions = append(wb.conditions, NewOrCondition(k, op, v))
	return wb
}

//
func (wb *criteriaBuilder) AndBlock(cns ...Condition) *criteriaBuilder {
	wb.conditions = append(wb.conditions, NewAndBlockCondition(cns...))
	return wb
}

func (wb *criteriaBuilder) OrBlock(cns ...Condition) *criteriaBuilder {
	wb.conditions = append(wb.conditions, NewOrBlockCondition(cns...))
	return wb
}

func (crt *criteriaBuilder) limitAndOffset(withLimit bool) string {
	builder := strings.Builder{}

	if crt.limit > 0 && crt.pageIndex > 0 && withLimit {
		// limit
		builder.WriteString(fmt.Sprintf("limit %d ", crt.limit))
		// offset
		builder.WriteString(fmt.Sprintf("offset %d ", (crt.pageIndex-1)*crt.limit))
	}
	return builder.String()
}

// create sql after the where condition
func (crt *criteriaBuilder) string(withLimit bool) string {
	// var limit string
	// var offset string
	var sort string
	var groupBy string
	// limit
	// if crt.limit > 0 && crt.pageIndex > 0 && withLimit {
	// 	limit = fmt.Sprintf("limit %d", crt.limit)
	// 	offset = fmt.Sprintf("offset %d", (crt.pageIndex-1)*crt.limit)
	// }
	// group by
	if len(crt.groupBy) > 0 {
		groupBy = fmt.Sprintf("group by %s", strings.Join(crt.groupBy, ", "))
	}
	// sort
	if crt.sort != "" && withLimit {
		sort = fmt.Sprintf("order by %s", crt.sort)
	}
	return fmt.Sprintf("%s %s %s", groupBy, sort, crt.limitAndOffset(withLimit))
}

// mount sql startint variable in 'startsIn', musst be greater or equal than 1
// and add or not limit and offse base on 'withLimit'
func (crt *criteriaBuilder) Build(startsIn int, withLimit bool) string {
	// create sql
	var sql bytes.Buffer
	seq := startsIn
	// values
	for _, v := range crt.conditions {
		sql.WriteString(v.Exp(seq))
		seq = seq + v.Count()
	}
	//
	after := crt.string(withLimit)
	//
	sql.WriteString(after)
	//
	return sql.String()
}

func (crt *criteriaBuilder) BuildWithCount(tmpl string, startsIn int) string {
	// create sql
	var sql bytes.Buffer
	seq := startsIn
	// values
	for _, v := range crt.conditions {
		sql.WriteString(v.Exp(seq))
		seq = seq + v.Count()
	}
	//
	after := crt.string(false)
	//
	sql.WriteString(after)
	//
	newSql := fmt.Sprintf(`
		with slc as (%s %s),
		count as (select count(*) from slc)
		select *, (select * from count) from slc
		%s
	`, tmpl, sql.String(), crt.limitAndOffset(true))
	return newSql
}

// return values from criteria and throw new values n
// on the beggining of the array case needed
func (crt *criteriaBuilder) Values(n ...any) []any {
	// create sql
	// values
	values := make([]any, 0)
	if n != nil {
		values = append(values, n...)
	}
	for _, v := range crt.conditions {
		values = append(values, v.Values()...)
	}
	//
	return values
}

// set page
func (b *criteriaBuilder) SetPage(page int) *criteriaBuilder {
	b.pageIndex = page
	return b
}

// set rows limit
func (b *criteriaBuilder) SetLimit(limit int) *criteriaBuilder {
	b.limit = limit
	return b
}

// add group by clause
func (b *criteriaBuilder) AddGroupBy(field string) *criteriaBuilder {
	b.groupBy = append(b.groupBy, field)
	return b
}

// set sort by
func (b *criteriaBuilder) SetSort(sort string) *criteriaBuilder {
	b.sort = sort
	return b
}

func (b *criteriaBuilder) Clone() Criteria {
	builder := *b
	return &builder
}

func NewCriteriaBuilder() Criteria {
	return &criteriaBuilder{
		pageIndex:    -1,
		limit:        10,
		sort:         "",
		groupBy:      []string{},
		filterValues: []any{},
		conditions:   make([]Condition, 0),
	}
}

func EmptyCriteria() Criteria {
	return &criteriaBuilder{
		pageIndex:    -1,
		limit:        -1,
		sort:         "",
		groupBy:      []string{},
		filterValues: []any{},
		conditions:   make([]Condition, 0),
	}
}
