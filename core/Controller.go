package core

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Code    int
	Message string
  Data interface{}
}

func (con Controller) Output(c *gin.Context) {
  fmt.Println(con)
	c.JSON(http.StatusOK, gin.H{
		"code":    con.Code,
		"message":    con.Message,
		"data":    con.Data,
	})
}
