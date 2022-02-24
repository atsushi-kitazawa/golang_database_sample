package postgres

import (
	"database/sql"
)

func Initdb() *sql.DB {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=15432 user=postgres password=mysecretpassword dbname=golang_db sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}
