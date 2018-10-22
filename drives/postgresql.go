package drives

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	SQL *sqlx.DB
}

var dbConn = &DB{}

func PostgreSQL() (*DB, error) {
	connStr := "postgres://postgres:123456@localhost/blog?sslmode=disable"
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		panic(err)
		return nil, err
	}

	dbConn.SQL = db

	return dbConn, nil
}
