package user

import (
	"gin_demo/core"
	"gin_demo/models/member"
	"github.com/gin-gonic/gin"
)
var con core.Controller

func Lists(c *gin.Context) {
  var mm member.Model;
  var page Page
  var where core.Where
  c.Bind(&page)
  con.Code = 200
  con.Message = "success"
  where := map[string]interfaces{id:1}
  con.Data.Total = mm.Lists(page, where)
  con.Data.Total = mm.Lists(page, where)
	con.Output(c)
}
