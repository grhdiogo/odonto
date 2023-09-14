package pgclient

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"odonto/internal/infra/criteria"

	"odonto/internal/domain/person"
)

var (
	ErrSavePerson   = errors.New("error on save person")
	ErrFindPerson   = errors.New("error on find person")
	ErrListPerson   = errors.New("error on list person")
	ErrUpdatePerson = errors.New("error on update person")
	ErrRemovePerson = errors.New("error on remove person")
)

type personRepositoryImpl struct {
	ctx context.Context
	tx  *sql.Tx
}

func (r *personRepositoryImpl) scanIterator(rows *sql.Rows) (*person.Entity, error) {
	//
	pid := sql.NullString{}
	name := sql.NullString{}
	cpf := sql.NullString{}
	email := sql.NullString{}
	birthdate := sql.NullTime{}
	// scan columns
	err := rows.Scan(
		&pid,
		&name,
		&cpf,
		&email,
		&birthdate,
	)
	if err != nil {
		return nil, err
	}
	// create entity
	entity := new(person.Entity)
	//
	if pid.Valid {
		entity.ID.Pid = pid.String
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
	return entity, nil
}

func (r *personRepositoryImpl) StoreAll(es []person.Entity) error {
	for _, e := range es {
		err := r.Store(&e)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *personRepositoryImpl) Store(e *person.Entity) error {
	// create sql
	sqlText := `
		insert into tb_person ( 
			pid, 
			name, 
			cpf, 
			email, 
			birthdate
		)
		values (
			$1 , $2 , $3 , $4 , $5 
		) `

	result, err := r.tx.Exec(sqlText,
		e.ID.Pid,
		e.Name,
		e.Cpf,
		e.Email,
		e.Birthdate,
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
		return ErrSavePerson
	}
	// result
	return nil
}

func (r *personRepositoryImpl) Update(e *person.Entity) error {
	// sql text
	sqlText := ` 
	update tb_person set 
		name = $1,
		cpf = $2,
		email = $3,
		birthdate = $4,
		updated_at = CURRENT_TIMESTAMP
		where 
			pid = $5 
	`

	result, err := r.tx.Exec(sqlText,
		e.Name,
		e.Cpf,
		e.Email,
		e.Birthdate,
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
		return ErrUpdatePerson
	}
	// result
	return nil
}

func (r *personRepositoryImpl) Remove(id *person.Identity) error {
	// sql text
	sqlText := `
	update tb_person set
		deleted_at = CURRENT_TIMESTAMP
	where
		deleted_at is null
		and pid = $1
	`

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
		return ErrRemovePerson
	}
	// result
	return nil
}

func (r *personRepositoryImpl) Find(crt criteria.Criteria) (*person.Entity, error) {
	// sql text
	sqlText := fmt.Sprintf(`
		select 
			pid,
			name,
			cpf,
			email,
			birthdate
			from tb_person
			where deleted_at is null
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
	return nil, ErrFindPerson
}

func (r *personRepositoryImpl) List(crt criteria.Criteria) ([]person.Entity, error) {
	// sql text
	sqlText := fmt.Sprintf(`
		select 
			pid, 
			name, 
			cpf, 
			email, 
			birthdate
		from tb_person
		where deleted_at is null
		%s
	`, crt.Build(1, false))

	// execute query in database
	rows, err := r.tx.Query(sqlText, crt.Values()...)
	if err != nil {
		return nil, err
	}
	// iterate
	result := make([]person.Entity, 0)
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

func NewPersonRepository(ctx context.Context, tx *sql.Tx) person.Repository {
	return &personRepositoryImpl{
		ctx: ctx,
		tx:  tx,
	}
}
