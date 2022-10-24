package core

import (
	"math"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Page     int
	Offset   int
	PageSize int
	Total    int
	Type     string
}

func (p *Pagination) Init(c *gin.Context, total int) {
	c.Bind(&p)
	p.Total = total
  p.CreateLimit()
}

func (p *Pagination) CreateLimit() {
	if p.PageSize < 1 {
		p.PageSize = 10
	}
	total_page := int(math.Ceil(float64(p.Total / p.PageSize)))
	if p.Page > total_page {
		p.Page = total_page
	}
	if p.Page < 1 {
		p.Page = 1
	}
	p.Offset = (p.Page - 1) * p.PageSize
}
