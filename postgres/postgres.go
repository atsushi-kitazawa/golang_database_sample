package postgres

import (
	"database/sql"
	"fmt"
)

func Initdb() *sql.DB {
	DBMS := "postgres"
	USER := "postgres"
	PASS := "password"
	HOST := "my_postgres_14"
	PORT := 5432
	DBNAME := "golang_db"

	// CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	CONNECT := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, USER, PASS, DBNAME)
	db, err := sql.Open(DBMS, CONNECT)
	if err != nil {
		panic(err)
	}
	return db
}
