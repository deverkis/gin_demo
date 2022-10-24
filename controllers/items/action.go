package user

import (
	"fmt"
	"gin_demo/core"
	"gin_demo/models/items"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Keyword struct {
	Category string
}

func Lists(c *gin.Context) {
	var model items.ItemsDao
	var pagination core.Pagination
	var keyword Keyword
	c.Bind(&keyword)
	total, err := model.ListsCount(keyword)
	pagination.Init(c, total)
	list, err := model.Lists(pagination, keyword)
	if err != nil {
		fmt.Println("err", err)
	}
	var Cc = core.BaseAction{Code: 200, Message: "success"}
	Cc.Lists(pagination, list)
	c.JSON(http.StatusOK, Cc)
}
