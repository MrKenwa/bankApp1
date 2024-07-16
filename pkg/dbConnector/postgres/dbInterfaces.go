package postgres

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type DBOps interface {
	Beginx() (*sqlx.Tx, error)
	Select(dest interface{}, query string, args ...interface{}) error
	Get(dest interface{}, query string, args ...interface{}) error
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
	Close() error
}

type PostgresDB struct {
	DB *sqlx.DB
}

func (d *PostgresDB) Beginx() (*sqlx.Tx, error) {
	return d.DB.Beginx()
}

func (d *PostgresDB) Select(dest interface{}, query string, args ...interface{}) error {
	return d.DB.Select(dest, query, args...)
}

func (d *PostgresDB) Get(dest interface{}, query string, args ...interface{}) error {
	return d.DB.Get(dest, query, args...)
}

func (d *PostgresDB) QueryRow(query string, args ...interface{}) *sql.Row {
	return d.DB.QueryRow(query, args...)
}

func (d *PostgresDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return d.DB.Exec(query, args...)
}

func (d *PostgresDB) Close() error {
	return d.DB.Close()
}
