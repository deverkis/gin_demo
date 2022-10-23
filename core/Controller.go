package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
type Controller struct {
	Code    int
	Message string
  Data interface{}
}

type DataVo struct {
	Total    int
	PageSize    int
	Page    int
  List interface{}
}

func (con Controller) Output(page Pagination, list interface{}) {
  data:=DataVo{
    Total : page.Total,
    PageSize : page.PageSize,
    Page : page.Page,
    List: list,
  }
	c.JSON(http.StatusOK, data)
}
