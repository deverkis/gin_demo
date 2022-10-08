package core

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Code    int
	Message string
}

func (con Controller) Success(c *gin.Context, data interface{}) {
	fmt.Println(data)
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": con.Message,
		"data":    data,
	})
}
