package member

import (
  "fmt"
    "database/sql"
	"gin_demo/core"
)

type Model struct {
  core.Model
}

type Items struct {
  Id int64
  Ename string 
  Name string 
}

func (m Model) Lists() (Items, error) {
  m.Table = "items"
  var alb Items
  id := 1
  row := m.FetchRow(id)
  if err := row.Scan(&alb.Id, &alb.Ename, &alb.Name); err != nil {
      if err == sql.ErrNoRows {
          return alb, fmt.Errorf("albumsById %d: no such album", id)
      }
      return alb, fmt.Errorf("albumsById %d: %v", id, err)
  }
  return alb, nil
}