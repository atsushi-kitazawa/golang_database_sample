package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/atsushi-kitazawa/golang_database_sample/postgres"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type user struct {
	Id   int    `json:Id`
	Name string `json:Name`
}

func main() {
	doMain()
}

func doMain() {
	r := gin.Default()
	// get user list
	r.GET("/postgres/users", func(c *gin.Context) {
		db := postgres.Initdb()
		db.Query("SELECT * FROM users")
		rows, err := db.Query("SELECT * FROM users")
		if err != nil {
			log.Fatalln("err=", err)
			return
		}

		var users []user
		for rows.Next() {
			t := user{}
			rows.Scan(&t.Id, &t.Name)
			users = append(users, t)
		}
		fmt.Println(users)
		c.IndentedJSON(http.StatusOK, users)
	})

	// create user
	r.POST("/postgres/users", func(c *gin.Context) {
		var u user
		if err := c.ShouldBind(&u); err != nil {
			c.String(http.StatusBadRequest, "bad request")
			return
		}
		db := postgres.Initdb()
		tx, err := db.Begin()
		if err != nil {
			log.Fatalln("err=", err)
			return
		}
		_, err = tx.Exec("INSERT INTO users VALUES ($1, $2)", u.Id, u.Name)
		if err != nil {
			log.Fatalln("err=", err)
			tx.Rollback()
			return
		}
		if err = tx.Commit(); err != nil {
			log.Fatalln("err=", err)
			return
		}
	})
	r.Run(":18080")
}
