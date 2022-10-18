package member

import (
	"gin_demo/core"
)

type Model struct {
  core.Model
}

func (m Model) Lists() string {
  m.table = "member"
  return m.FetchRow()
}