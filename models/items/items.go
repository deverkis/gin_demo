package items

import (
	"fmt"
	"gin_demo/core"
	"gin_demo/models"
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

func (c ItemsDao) ListsCount(where interface{}) (ret int, err error) {
	var total int64
	fmt.Println("ListsCount ")
	db := models.DB.Table("items")
	db.Count(&total)
	fmt.Println("total", total)
	ret = int(total)
	return
}

func (c ItemsDao) Lists(page core.Pagination, where interface{}) (vos []ListVO, err error) {
	//var vos []ListVO
	db := models.DB.Table("items")
	db.Select("id,ename,name,price")
	fmt.Println(page)
	db.Offset(page.Offset).Limit(page.PageSize).Find(&vos)
	return
}
