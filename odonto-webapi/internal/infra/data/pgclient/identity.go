package pgclient

import (
	"context"
	"database/sql"
	"sync"
)

//NewIdentity create or update identity for an entity
func NewIdentity(ctx context.Context, tx *sql.Tx, entity string) (int64, error) {
	// create a locker
	locker := sync.Mutex{}
	locker.Lock()
	defer locker.Unlock()
	// sql for select
	selectSQLText := `select value from identity where entity = $1`
	insertSQLText := `insert into identity (entity, value) values ($1, $2)`
	updateSQLText := `update identity set value = $2 where entity = $1`
	// recover
	rows, err := tx.Query(selectSQLText, entity)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	//
	identity := int64(1)
	if rows.Next() {
		// retrieve
		err = rows.Scan(&identity)
		if err != nil {
			return 0, err
		}
		// increment identity value
		identity = identity + 1
		// update
		_, err = tx.Exec(updateSQLText, entity, identity)
		if err != nil {
			return 0, err
		}
	} else {
		_, err = tx.Exec(insertSQLText, entity, identity)
		if err != nil {
			return 0, err
		}
	}
	//
	return identity, nil
}
