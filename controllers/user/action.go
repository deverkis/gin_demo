package user

import (
	"gin_demo/core"
	"gin_demo/models/member"
	"github.com/gin-gonic/gin"
)
var con core.Controller

func Lists(c *gin.Context) {
  var mm member.Model;
  con.Code = 200
  con.Message = "success"
  con.Data = mm.Lists()
	con.Output(c)
}
