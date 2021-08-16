package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

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
	return
    }

    // select table
    rows, err := db.Query("SELECT * FROM t1")
    if err != nil {
	log.Fatalln("err=", err)
	return
    }

    for rows.Next() {
	t := t1{}
	rows.Scan(&t.id, &t.name)
	fmt.Println(t.id, t.name)
    }

    // insert table
    tableCount, err := db.Query("SELECT COUNT(*) FROM t1")
    if err != nil {
	log.Fatalln("err=", err)
	return
    }
    var cnt int
    tableCount.Next()
    tableCount.Scan(&cnt)

    tx, err := db.Begin()
    if err != nil {
	log.Fatalln("err=", err)
	return
    }
    id := cnt + 1
    name := "aaa" + strconv.Itoa(id)
    _, err = tx.Exec("INSERT INTO t1 VALUES ($1, $2)", id, name)
    if err != nil {
	log.Fatalln("err=", err)
	tx.Rollback()
	return
    }
    if err = tx.Commit(); err != nil {
	log.Fatalln("err=", err)
	return
    }
}
