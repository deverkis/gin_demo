package core

import (
  "math"
)

type Pagination struct {
	Page    int
  Offset int
	PageSize int
  Total int
  Type string
}

func (p *Pagination) Create(){
  if p.PageSize < 1 {
    p.PageSize = 10
  }
  total_page := int(math.Ceil(float64(p.Total / p.PageSize)));
  if p.Page > total_page{
			p.Page = total_page;
  }
  if p.Page < 1{
    p.Page = 1;
  }
  p.Offset= (p.Page-1) * p.PageSize
}
