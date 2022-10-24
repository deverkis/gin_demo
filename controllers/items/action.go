package user

import (
	"fmt"
	"gin_demo/core"
	"gin_demo/models/items"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Lists(c *gin.Context) {
	var model items.ItemsDao
	var pagination core.Pagination
	total,_, err := model.Lists(core.Pagination{Type:"count"}, c)
	pagination.Init(c, total)
	_, list, err := model.Lists(pagination, c)
	if err != nil {
		fmt.Println("err", err)
	}
	var Cc = core.BaseAction{Code: 200, Message: "success"}
	Cc.Lists(pagination, list)
	c.JSON(http.StatusOK, Cc)
}
