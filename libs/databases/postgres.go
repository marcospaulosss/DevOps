package databases

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	log "backend/libs/logger"
)

type Postgres struct {
	Connection *sqlx.DB
	ConnStr    string
	Tx         *sqlx.Tx
}

func NewPostgres(connStr string) Database {
	db := &Postgres{}
	db.ConnStr = connStr
	return db
}

func (db *Postgres) Connect() (*sqlx.DB, error) {
	var err error
	if db.Connection == nil {
		db.Connection, err = db.getConnection()
	}
	return db.Connection, err
}

func (db *Postgres) GetConnection() *sqlx.DB {
	conn, err := db.Connect()
	if err != nil {
		log.Error("Falhou ao conectar com o banco de dados", err.Error())
	}
	return conn
}

func (db *Postgres) getConnection() (*sqlx.DB, error) {
	conn, err := sqlx.Connect("postgres", db.ConnStr)
	if err != nil {
		log.Error("Failed connecting to the database:", db.ConnStr, err.Error())
		return conn, err
	}
	conn.SetMaxOpenConns(20)
	conn.SetMaxIdleConns(2)
	conn.SetConnMaxLifetime(time.Nanosecond)
	return conn, err
}

func (db *Postgres) Begin() (*sqlx.Tx, error) {
	conn := db.GetConnection()

	tx, err := conn.Beginx()
	db.Tx = tx

	return tx, err
}

func (db *Postgres) GetTx() *sqlx.Tx {
	return db.Tx
}

func (db *Postgres) Rollback() error {
	return db.Tx.Rollback()
}

func (db *Postgres) Commit() error {
	return db.Tx.Commit()
}
