package items

import (
	"fmt"
	"gin_demo/core"
	"gin_demo/models"
	"github.com/gin-gonic/gin"
)

type ItemsDao struct {
	ID    uint
	Ename string
	Name  string
	Price int
}
type ListVO struct {
	ID    uint   `json:"id"`
	Ename string `json:"ename"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func (ItemsDao) TableName() string {
	return "items"
}

type ListsKeyword struct {
	Category string `json:"category" form:"category"`
}

func (m ItemsDao) Lists(page core.Pagination, c *gin.Context) (total int64, vos []ListVO, err error) {
	//var vos []ListVO
  var keyword ListsKeyword
	db := models.DB.Table("items")
  c.Bind(&keyword)
  //where
  if keyword.Category != "" {
    db.Where("category",keyword.Category)
  }
  if(page.Type == "count") {
    db.Count(&total)
    return
  }
	db.Select("id,ename,name,price")
	fmt.Println(page)
	db.Offset(page.Offset).Limit(page.PageSize).Find(&vos)
	return
}
