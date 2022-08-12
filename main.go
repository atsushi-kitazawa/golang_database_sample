package main

import (
	"net/http"

	"github.com/atsushi-kitazawa/golang_database_sample/database"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type errorResponse struct {
	msg string
}

func main() {
	doMain()
}

func doMain() {
	r := gin.Default()
	// get user list
	r.GET("/postgres/users", func(c *gin.Context) {
		users, err := database.GetUsers()
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, errorResponse{
				msg: err.Error(),
			})
			return
		}
		c.IndentedJSON(http.StatusOK, users)
	})

	// create user
	r.POST("/postgres/users", func(c *gin.Context) {
		user := database.NewEmptyUser()
		if err := c.ShouldBind(user); err != nil {
			c.IndentedJSON(http.StatusBadRequest, errorResponse{
				msg: "bad request",
			})
			return
		}
		err := database.CreateUser(user)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, errorResponse{
				msg: err.Error(),
			})
			return
		}
	})
	r.Run(":18080")
}
