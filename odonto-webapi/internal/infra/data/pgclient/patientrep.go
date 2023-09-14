package pgclient

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"odonto/internal/infra/criteria"

	"odonto/internal/domain/patient"
)

var (
	ErrSavePatient   = errors.New("error on save patient")
	ErrFindPatient   = errors.New("error on find patient")
	ErrListPatient   = errors.New("error on list patient")
	ErrUpdatePatient = errors.New("error on update patient")
	ErrRemovePatient = errors.New("error on remove patient")
)

type patientRepositoryImpl struct {
	ctx context.Context
	tx  *sql.Tx
}

func (r *patientRepositoryImpl) scanIteratorGeneral(rows *sql.Rows, wc bool) (*patient.Entity, int, error) {
	//
	personPid := sql.NullString{}
	name := sql.NullString{}
	cpf := sql.NullString{}
	email := sql.NullString{}
	birthdate := sql.NullTime{}
	count := sql.NullInt64{}

	vars := []any{
		&personPid,
		&name,
		&cpf,
		&email,
		&birthdate,
	}
	if wc {
		vars = append(vars, &count)
	}
	// scan columns
	err := rows.Scan(
		vars...,
	)
	if err != nil {
		return nil, 0, err
	}
	// create entity
	entity := new(patient.Entity)
	//
	if personPid.Valid {
		entity.ID.PersonPid = personPid.String
	}
	if name.Valid {
		entity.Name = name.String
	}
	if cpf.Valid {
		entity.Cpf = cpf.String
	}
	if email.Valid {
		entity.Email = email.String
	}
	if birthdate.Valid {
		entity.Birthdate = birthdate.Time
	}
	// result
	return entity, int(count.Int64), nil
}

func (r *patientRepositoryImpl) scanIterator(rows *sql.Rows) (*patient.Entity, error) {
	ent, _, err := r.scanIteratorGeneral(rows, false)
	return ent, err
}

func (r *patientRepositoryImpl) StoreAll(es []patient.Entity) error {
	for _, e := range es {
		err := r.Store(&e)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *patientRepositoryImpl) Store(e *patient.Entity) error {
	// create sql
	sqlText := `
		insert into tb_patient ( 
			person_pid
		)
		values (
			$1 
		) `

	result, err := r.tx.Exec(sqlText,
		e.ID.PersonPid,
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
		return ErrSavePatient
	}
	// result
	return nil
}

func (r *patientRepositoryImpl) Update(e *patient.Entity) error {
	// sql text
	sqlText := ` 
	update tb_patient set 
		updated_at = CURRENT_TIMESTAMP
		where 
			person_pid = $1 
	`

	result, err := r.tx.Exec(sqlText,
		e.ID.PersonPid,
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
		return ErrUpdatePatient
	}
	// result
	return nil
}

func (r *patientRepositoryImpl) Remove(id *patient.Identity) error {
	// sql text
	sqlText := `
	update tb_patient set
		deleted_at = CURRENT_TIMESTAMP
	where
		deleted_at is null
		and person_pid = $1
	`

	result, err := r.tx.Exec(sqlText,
		id.PersonPid,
	)
	if err != nil {
		return err
	}
	aff, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if aff != 1 {
		return ErrRemovePatient
	}
	// result
	return nil
}

func (r *patientRepositoryImpl) Find(crt criteria.Criteria) (*patient.Entity, error) {
	// sql text
	sqlText := fmt.Sprintf(`
		select 
			patient.person_pid,
			person.name,
			person.cpf,
			person.email,
			person.birthdate
			from tb_patient patient
			inner join tb_person person on patient.person_pid = person.pid
			where patient.deleted_at is null
			%s
			`, crt.Build(1, false))

	// execute query in database
	rows, err := r.tx.Query(sqlText, crt.Values()...)
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
	return nil, ErrFindPatient
}

func (r *patientRepositoryImpl) List(crt criteria.Criteria) ([]patient.Entity, int, error) {
	// sql text
	sqlText := `
		select 
			patient.person_pid,
			person.name,
			person.cpf,
			person.email,
			person.birthdate
		from tb_patient patient
		inner join tb_person person on patient.person_pid = person.pid
		where patient.deleted_at is null
	`
	// execute query in database
	rows, err := r.tx.Query(crt.BuildWithCount(sqlText, 1), crt.Values()...)
	if err != nil {
		return nil, 0, err
	}
	// iterate
	result := make([]patient.Entity, 0)
	var count int
	//
	defer rows.Close()
	for rows.Next() {
		value, c, err := r.scanIteratorGeneral(rows, true)
		if err != nil {
			return nil, 0, err
		}
		result = append(result, *value)
		count = c
	}
	// success
	return result, count, nil
}

func NewPatientRepository(ctx context.Context, tx *sql.Tx) patient.Repository {
	return &patientRepositoryImpl{
		ctx: ctx,
		tx:  tx,
	}
}
