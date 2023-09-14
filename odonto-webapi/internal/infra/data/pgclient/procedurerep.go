package pgclient

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"odonto/internal/infra/criteria"

	"odonto/internal/domain/procedure"
)

var (
	ErrSaveProcedure   = errors.New("error on save procedure")
	ErrFindProcedure   = errors.New("error on find procedure")
	ErrListProcedure   = errors.New("error on list procedure")
	ErrUpdateProcedure = errors.New("error on update procedure")
	ErrRemoveProcedure = errors.New("error on remove procedure")
)

type procedureRepositoryImpl struct {
	ctx context.Context
	tx  *sql.Tx
}

func (r *procedureRepositoryImpl) scanIterator(rows *sql.Rows) (*procedure.Entity, error) {
	//
	pid := sql.NullString{}
	name := sql.NullString{}
	value := sql.NullFloat64{}
	// scan columns
	err := rows.Scan(
		&pid,
		&name,
		&value,
	)
	if err != nil {
		return nil, err
	}
	// create entity
	entity := new(procedure.Entity)
	//
	if pid.Valid {
		entity.ID.Pid = pid.String
	}
	if name.Valid {
		entity.Name = name.String
	}
	if value.Valid {
		entity.Value = value.Float64
	}
	// result
	return entity, nil
}

func (r *procedureRepositoryImpl) StoreAll(es []procedure.Entity) error {
	for _, e := range es {
		err := r.Store(&e)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *procedureRepositoryImpl) Store(e *procedure.Entity) error {
	// create sql
	sqlText := `
		insert into tb_procedure ( 
			pid, 
			name, 
			value
		)
		values (
			$1 , $2 , $3
		) `

	result, err := r.tx.Exec(sqlText,
		e.ID.Pid,
		e.Name,
		e.Value,
	)
	if err != nil {
		return err
	}
	// check if insert was affected
	aff, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if aff != 1 {
		return ErrSaveProcedure
	}
	// result
	return nil
}

func (r *procedureRepositoryImpl) Update(e *procedure.Entity) error {
	// sql text
	sqlText := ` 
	update tb_procedure set 
		name = $1,
		value = $2,
		updated_at = CURRENT_TIMESTAMP
		where 
			pid = $3 
	`

	result, err := r.tx.Exec(sqlText,
		e.Name,
		e.Value,
		e.ID.Pid,
	)
	if err != nil {
		return err
	}
	// check if update was affected
	aff, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if aff != 1 {
		return ErrUpdateProcedure
	}
	// result
	return nil
}

func (r *procedureRepositoryImpl) Remove(id *procedure.Identity) error {
	// sql text
	sqlText := `
	update tb_procedure set
		deleted_at = CURRENT_TIMESTAMP
	where
		deleted_at is null
		and pid = $1
	`
	fmt.Println(">>>>", id.Pid)
	result, err := r.tx.Exec(sqlText,
		id.Pid,
	)
	if err != nil {
		return err
	}
	// check if insert was affected
	aff, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if aff == 0 {
		return ErrRemoveProcedure
	}
	// result
	return nil
}

func (r *procedureRepositoryImpl) Find(id procedure.Identity) (*procedure.Entity, error) {
	// sql text
	sqlText := `
		select 
			pid,
			name,
			value
			from tb_procedure
			where pid = $1
			`

	// execute query in database
	rows, err := r.tx.Query(sqlText, id)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		ent, err := r.scanIterator(rows)
		if err != nil {
			rows.Close()
			return nil, err
		}
		rows.Close()
		return ent, nil
	}
	//retun
	return nil, ErrFindProcedure
}

func (r *procedureRepositoryImpl) List(crt criteria.Criteria) ([]procedure.Entity, error) {
	// sql text
	sqlText := fmt.Sprintf(`
		select 
			pid, 
			name, 
			value
		from tb_procedure
		where deleted_at is null
		%s
	`, crt.Build(1, false))

	// execute query in database
	rows, err := r.tx.Query(sqlText, crt.Values()...)
	if err != nil {
		return nil, err
	}
	// iterate
	result := make([]procedure.Entity, 0)
	for rows.Next() {
		value, err := r.scanIterator(rows)
		if err != nil {
			rows.Close()
			return nil, err
		}
		result = append(result, *value)
	}
	rows.Close()
	// success
	return result, nil
}

func NewProcedureRepository(ctx context.Context, tx *sql.Tx) procedure.Repository {
	return &procedureRepositoryImpl{
		ctx: ctx,
		tx:  tx,
	}
}
