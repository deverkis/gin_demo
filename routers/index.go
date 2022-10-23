package routers

import (
	items "gin_demo/controllers/items"
	"github.com/gin-gonic/gin"
)

func Index(r *gin.Engine) {
	r1 := r.Group("/items")
	{
		r1.GET("/lists", items.Lists)
	}



}
