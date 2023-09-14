package pgclient

// type authaccessRepositoryImpl struct {
// 	tx  sql.Tx
// 	ctx context.Context
// }

// func (r *authaccessRepositoryImpl) scanIterator(rows pgx.Rows) (*authaccess.Entity, error) {
// 	//
// 	token := sql.NullString{}
// 	userUid := sql.NullString{}
// 	expiredAt := sql.NullTime{}
// 	// scan columns
// 	err := rows.Scan(
// 		&token,
// 		&userUid,
// 		&expiredAt,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// create entity
// 	entity := new(authaccess.Entity)
// 	//
// 	if token.Valid {
// 		entity.ID.Token = token.String
// 	}
// 	if userUid.Valid {
// 		entity.UserUid = userUid.String
// 	}
// 	if expiredAt.Valid {
// 		entity.ExpiredAt = expiredAt.Time
// 	}
// 	return entity, nil
// }

// func (r *authaccessRepositoryImpl) Store(e *authaccess.Entity) error {
// 	// create sql
// 	sqlText := `
// 		insert into tb_auth_access (
// 			token,
// 			user_uid,
// 			expired_at
// 		)
// 		values (
// 			$1, $2, $3
// 		) `

// 	result, err := r.tx.Exec(r.ctx,
// 		sqlText, e.ID.Token, e.UserUid, e.ExpiredAt,
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	// affected rows
// 	affected := result.RowsAffected()

// 	// check if insert was affected
// 	if affected != 1 {
// 		return errors.New("error on save authaccess")
// 	}
// 	// result
// 	return nil
// }

// func (r *authaccessRepositoryImpl) Remove(id *authaccess.Identity) error {
// 	// sql text
// 	sqlText := `
// 	update tb_auth_access set
// 		deleted_at = now()
// 	where
// 	token = $1 and deleted_at is null

// 	`

// 	result, err := r.tx.Exec(r.ctx,
// 		sqlText, id.Token,
// 	)
// 	if err != nil {
// 		return err
// 	}
// 	// affected rows
// 	affected := result.RowsAffected()
// 	// check if insert was affected
// 	if affected == 0 {
// 		return errors.New("error on delete authaccess")
// 	}
// 	// result
// 	return nil
// }

// func (r *authaccessRepositoryImpl) Find(id *authaccess.Identity) (*authaccess.Entity, error) {
// 	// sql text
// 	sqlText := `
// 		select
// 			token,
// 			user_uid,
// 			expired_at
// 			from tb_auth_access
// 		where deleted_at is null
// 	 		and token = $1
// 			and expired_at > now()
// 	`

// 	// execute query in database
// 	rows, err := r.tx.Query(r.ctx,
// 		sqlText, id.Token,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	if rows.Next() {
// 		return r.scanIterator(rows)
// 	}
// 	//retun
// 	return nil, errors.New("authaccess not found")
// }

// func NewAuthAccessRepository(ctx context.Context, tx pgx.Tx, tenant string) authaccess.Repository {
// 	return &authaccessRepositoryImpl{
// 		tx:  tx,
// 		ctx: ctx,
// 	}
// }
