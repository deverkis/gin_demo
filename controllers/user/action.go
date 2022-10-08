package user

import (
	"fmt"
	"gin_demo/core"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	core.Controller
}

func (con Controller) Lists(c *gin.Context) {
	var data struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	data.Id = 123
	data.Name = "Hey"
	con.Message = "success"
	fmt.Println(data)
	con.Success(c, data)
}
