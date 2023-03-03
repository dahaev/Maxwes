package main

import (
	"maxwes_group/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	id = "id"
)

func main() {

	r := gin.Default()
	r.GET("/get-items/", GetItems)
	r.Run(":8080")
}

func GetItems(c *gin.Context) {

	number := c.QueryArray(id)

	if len(number) == 0 {

		c.JSON(http.StatusBadRequest, "bad url")
		return
	}

	answer := service.MembershipRequest(number)

	c.JSON(http.StatusOK, answer)
}
