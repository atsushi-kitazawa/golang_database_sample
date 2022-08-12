package database

import (
	"database/sql"
	"fmt"
	"log"
)

type user struct {
	Id   int    `json:"Id"`
	Name string `json:"Name"`
}

func NewEmptyUser() *user {
	return &user{}
}

func initdb() *sql.DB {
	DBMS := "postgres"
	USER := "postgres"
	PASS := "password"
	HOST := "my_postgres_14"
	PORT := 5432
	DBNAME := "postgres"

	// CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	CONNECT := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, USER, PASS, DBNAME)
	db, err := sql.Open(DBMS, CONNECT)
	if err != nil {
		panic(err)
	}
	return db
}

func GetUsers() (*[]user, error) {
	db := initdb()
	db.Query("SELECT * FROM users")
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Println("GetUsers() err=", err)
		return nil, err
	}

	var users []user
	for rows.Next() {
		t := user{}
		rows.Scan(&t.Id, &t.Name)
		users = append(users, t)
	}
	return &users, nil
}

func CreateUser(u *user) error {
	db := initdb()
	tx, err := db.Begin()
	if err != nil {
		log.Println("CreateUser() - Begin() err=", err)
		return err
	}
	_, err = tx.Exec("INSERT INTO users VALUES ($1, $2)", u.Id, u.Name)
	if err != nil {
		log.Println("CreateUser() - Exec() err=", err)
		tx.Rollback()
		return err
	}
	if err = tx.Commit(); err != nil {
		log.Println("CreateUser() - Commit() err=", err)
		return err
	}
	return nil
}
