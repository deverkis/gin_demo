package user

import (
  "fmt"
	"net/http"
	"gin_demo/core"
	"gin_demo/models/items"
	"github.com/gin-gonic/gin"
)
var cc core.Controller

type Keyword struct{
  Category string
}

type DataVo struct {
	Total    int
	PageSize    int
	Page    int
  List interface{}
}

func Lists(c *gin.Context) {
  var model items.ItemsDao
  var pagination core.Pagination
  var keyword Keyword
  c.Bind(&keyword)
  c.Bind(&pagination)
  total,err := model.ListsCount(keyword)
  pagination.Total = total
  pagination.Create()
  list,err := model.Lists(pagination,keyword)
  if err != nil {
    fmt.Println("err",err)
  }
  fmt.Println(pagination,list)
  data:=DataVo{
    Total : pagination.Total,
    PageSize : pagination.PageSize,
    Page : pagination.Page,
    List: list,
  }
	c.JSON(http.StatusOK, data) 
}
