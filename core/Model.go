package core

import (
  "fmt"
  "log"  
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)
var DB *gorm.DB

func init(){
  dsn := "root:123456@tcp(127.0.0.1:3306)/user?charset=utf8mb4&parseTime=True&loc=Local"
  DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }
}

type Where interface {
  
}
type Model struct {
  Table string
}

func (m Model) fetch_count() {

}

func (m Model) FetchRow(id int) (*sql.Row)  {
    m.connect()
    sql := fmt.Sprintf("SELECT id,ename,name FROM `%s` WHERE `id` = '%d'", m.Table, id)
    row := DB.QueryRow(sql)
    return row
}
