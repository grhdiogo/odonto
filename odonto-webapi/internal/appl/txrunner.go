package appl

import (
	"database/sql"

	"odonto/internal/infra/data/pgclient"
	"odonto/internal/infra/fails"
)

// timeout in seconds
const txTimeout = 1000

type TxRunnerFunc func(tx *sql.Tx) (interface{}, error)

func TxRun(txf TxRunnerFunc) (interface{}, error) {
	// get pool of connections
	connManager := pgclient.GetManager()
	poolConn, err := connManager.GetConn()
	if err != nil {
		return "", fails.ProcessStack(100, ErrGetConnection, err)
	}
	//
	//create transaction repository
	tx, err := poolConn.StartTx()
	if err != nil {
		return "", fails.ProcessStack(101, ErrStartTransaction, err)
	}
	defer tx.Rollback()
	// operation
	result, err := txf(tx)
	if err != nil {
		return "", fails.ProcessStack(102, ErrStartTransaction, err)
	}
	// result
	return result, tx.Commit()
}
