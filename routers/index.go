package routers

import (
	"gin_demo/controllers/user"

	"github.com/gin-gonic/gin"
)

func Index(r *gin.Engine) {
	router_user := r.Group("/user")
	{
		router_user.GET("/lists", user.Lists)
	}



}
