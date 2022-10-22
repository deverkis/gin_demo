package member

import (
  "fmt"
	"gin_demo/core"
  "gorm.io/gorm"
  "gorm.io/driver/mysql"
)

type ApiMember struct {
  ID   uint
  Name string
}

func Lists(page core.Page, where map[string]interface{}) {
  model :=core.DB.Table("member")
  if(page.Type=='count'){
    var total int
    model.Count(@page.Total)
    return total
  } 
  var apiMember []ApiMember
  model.Select("id,created_at,updated_at,name,age,birth")
  model.Offet((page.Page-) * page.PageSize).Limit(page.PageSize)
  model.Find(&apiMember)
  return apiMember
}