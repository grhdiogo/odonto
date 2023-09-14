package pgclient

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"odonto/internal/domain/user"
	"odonto/internal/infra/criteria"

	"github.com/jackc/pgx/v4"
)

type userRepositoryImpl struct {
	tx  pgx.Tx
	ctx context.Context
}

var (
	ErrUserNotFound = errors.New("user not found")
)

func (r *userRepositoryImpl) scanIterator(rows pgx.Rows) (*user.Entity, error) {
	//
	personPid := sql.NullString{}
	kind := sql.NullString{}
	identifier := sql.NullString{}
	secret := sql.NullString{}
	// scan columns
	err := rows.Scan(
		&personPid,
		&kind,
		&identifier,
		&secret,
	)
	if err != nil {
		return nil, err
	}
	// create entity
	entity := new(user.Entity)
	//
	if personPid.Valid {
		entity.ID.PersonPID = personPid.String
	}
	if kind.Valid {
		entity.Kind = user.UserKind(kind.String)
	}
	if identifier.Valid {
		entity.Identifier = identifier.String
	}
	if secret.Valid {
		entity.Secret = secret.String
	}

	return entity, nil
}

func (r *userRepositoryImpl) Store(e *user.Entity) error {
	// create sql
	sqlText := `
		insert into tb_user ( 
			person_pid,
			kind,
			identifier,
			secret
		)
		values ($1, $2, $3, $4)
		`

	result, err := r.tx.Exec(r.ctx,
		sqlText, e.ID.PersonPID, e.Kind, e.Identifier, e.Secret,
	)
	if err != nil {
		return err
	}
	// check if insert was affected
	if result.RowsAffected() != 1 {
		return errors.New("error on save user")
	}
	// result
	return nil
}

func (r *userRepositoryImpl) Update(e *user.Entity) error {
	// sql text
	sqlText := ` update tb_user set
		secret = $3,
		updated_at = CURRENT_TIMESTAMP
		where 
			person_pid = $1
	`

	result, err := r.tx.Exec(r.ctx,
		sqlText,
		e.ID.PersonPID,
	)
	if err != nil {
		return err
	}
	// check if update was affected
	if result.RowsAffected() != 1 {
		return errors.New("error on update user")
	}
	// result
	return nil
}

func (r *userRepositoryImpl) Remove(id *user.Identity) error {
	// sql text
	sqlText := `
	update tb_user set
		deleted_at = CURRENT_TIMESTAMP
	where 
	person_pid = $1
	`

	result, err := r.tx.Exec(r.ctx,
		sqlText, id.PersonPID,
	)
	if err != nil {
		return err
	}
	// check if row was affected
	if result.RowsAffected() == 0 {
		return errors.New("error on delete user")
	}
	// result
	return nil
}

func (r *userRepositoryImpl) Find(id *user.Identity) (*user.Entity, error) {
	// sql text
	sqlText := `
		select 
			person_pid,
			kind,
			identifier,
			secret
			from tb_user
		where deleted_at is null
	 		and person_pid = $1 
	`
	// execute query in database
	rows, err := r.tx.Query(r.ctx,
		sqlText, id.PersonPID,
	)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		ent, err := r.scanIterator(rows)
		if err != nil {
			return nil, err
		}
		rows.Close()
		return ent, nil
	}
	//retun
	return nil, ErrUserNotFound
}

func (r *userRepositoryImpl) FindByIdentifierAndSecret(i, s string) (*user.Entity, error) {
	// sql text
	sqlText := `
		select 
			person_pid,
			kind,
			identifier,
			secret
			from tb_user
		where deleted_at is null
	 		and identifier = $1 
	 		and secret = $2 
	`
	// execute query in database
	rows, err := r.tx.Query(r.ctx,
		sqlText, i, s,
	)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		ent, err := r.scanIterator(rows)
		if err != nil {
			return nil, err
		}
		rows.Close()
		return ent, nil
	}
	//retun
	return nil, ErrUserNotFound
}

func (r *userRepositoryImpl) List(crt criteria.Criteria) ([]user.Entity, error) {
	// sql text
	sqlText := fmt.Sprintf(`
		select 
			person_pid,
			kind,
			identifier,
			secret		
		from tb_user
		where deleted_at is null
		%s
	`, crt.Build(1, false))
	// execute query in database
	rows, err := r.tx.Query(r.ctx, sqlText, crt.Values()...)
	if err != nil {
		return nil, err
	}
	// close rows
	// iterate
	result := make([]user.Entity, 0)
	for rows.Next() {
		value, err := r.scanIterator(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, *value)
	}
	rows.Close()
	// success
	return result, nil
}

func NewUserRepository(ctx context.Context, tx pgx.Tx) user.Repository {
	return &userRepositoryImpl{
		tx:  tx,
		ctx: ctx,
	}
}
