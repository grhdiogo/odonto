package pgclient

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"odonto/internal/infra/criteria"

	"odonto/internal/domain/doctor"
)

var (
	ErrSaveDoctor   = errors.New("error on save doctor")
	ErrFindDoctor   = errors.New("error on find doctor")
	ErrListDoctor   = errors.New("error on list doctor")
	ErrUpdateDoctor = errors.New("error on update doctor")
	ErrRemoveDoctor = errors.New("error on remove doctor")
)

type doctorRepositoryImpl struct {
	ctx context.Context
	tx  *sql.Tx
}

func (r *doctorRepositoryImpl) scanIteratorGeneral(rows *sql.Rows, wc bool) (*doctor.Entity, int, error) {
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
	entity := new(doctor.Entity)
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

func (r *doctorRepositoryImpl) scanIterator(rows *sql.Rows) (*doctor.Entity, error) {
	ent, _, err := r.scanIteratorGeneral(rows, false)
	return ent, err
}

func (r *doctorRepositoryImpl) StoreAll(es []doctor.Entity) error {
	for _, e := range es {
		err := r.Store(&e)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *doctorRepositoryImpl) Store(e *doctor.Entity) error {
	// create sql
	sqlText := `
		insert into tb_doctor ( 
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
		return ErrSaveDoctor
	}
	// result
	return nil
}

func (r *doctorRepositoryImpl) Update(e *doctor.Entity) error {
	// sql text
	sqlText := ` 
	update tb_doctor set 
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
		return ErrUpdateDoctor
	}
	// result
	return nil
}

func (r *doctorRepositoryImpl) Remove(id *doctor.Identity) error {
	// sql text
	sqlText := `
	update tb_doctor set
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
		return ErrRemoveDoctor
	}
	// result
	return nil
}

func (r *doctorRepositoryImpl) Find(crt criteria.Criteria) (*doctor.Entity, error) {
	// sql text
	sqlText := fmt.Sprintf(`
		select 
			doctor.person_pid,
			person.name,
			person.cpf,
			person.email,
			person.birthdate
			from tb_doctor doctor
			inner join tb_person person on doctor.person_pid = person.pid
			where doctor.deleted_at is null
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
	return nil, ErrFindDoctor
}

func (r *doctorRepositoryImpl) List(crt criteria.Criteria) ([]doctor.Entity, int, error) {
	// sql text
	sqlText := `
		select 
			doctor.person_pid,
			person.name,
			person.cpf,
			person.email,
			person.birthdate
		from tb_doctor doctor
		inner join tb_person person on doctor.person_pid = person.pid
		where doctor.deleted_at is null
	`
	// execute query in database
	rows, err := r.tx.Query(crt.BuildWithCount(sqlText, 1), crt.Values()...)
	if err != nil {
		return nil, 0, err
	}
	// iterate
	result := make([]doctor.Entity, 0)
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

func NewDoctorRepository(ctx context.Context, tx *sql.Tx) doctor.Repository {
	return &doctorRepositoryImpl{
		ctx: ctx,
		tx:  tx,
	}
}
