package user

import (
	"gin_demo/core"
	"github.com/gin-gonic/gin"
)
var con core.Controller

func Lists(c *gin.Context) {
  con.Code = 200
  con.Message = "success"
  con.Data = "abc"
	con.Output(c)
}
