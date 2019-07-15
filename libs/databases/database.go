package databases

import (
	"github.com/jmoiron/sqlx"
)

type Database interface {
	Connect() (*sqlx.DB, error)
	GetConnection() *sqlx.DB
	Begin() (*sqlx.Tx, error)
	GetTx() *sqlx.Tx
	Rollback() error
	Commit() error
}
