package core

import (
	"math"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Page     int64
	Offset   int
	PageSize int
	Total    int64
	Type     string
}

func (p *Pagination) Init(c *gin.Context, total int64) {
	c.Bind(&p)
	p.Total = total
  p.CreateLimit()
}

func (p *Pagination) CreateLimit() {
	if p.PageSize < 1 {
		p.PageSize = 10
	}
	total_page := int64(math.Ceil(float64(p.Total / int64(p.PageSize))))
	if p.Page > total_page {
		p.Page = total_page
	}
	if p.Page < 1 {
		p.Page = 1
	}
	p.Offset = int(p.Page - 1) * p.PageSize
}
