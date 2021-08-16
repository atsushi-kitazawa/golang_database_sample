package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type t1 struct {
    id int
    name string
}

func main() {
    fmt.Println("start database sample.")

    db, err := sql.Open("postgres", "host=127.0.0.1 port=15432 user=postgres dbname=postgres sslmode=disable")
    if err != nil {
	log.Fatalln("err=", err)
    }

    // select table
    rows, err := db.Query("SELECT * FROM t1")
    if err != nil {
	log.Fatalln("err=", err)
    }

    for rows.Next() {
	var t t1
	rows.Scan(&t.id, &t.name)
	fmt.Println(t.id)
	fmt.Println(t.name)
    }
}
