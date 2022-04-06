package handler

import (
	"email-template-generator/database"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	test := database.New()
	fmt.Println(test)
	c.String(http.StatusOK, "Hello World")
}
