package main

import (
	"gin_demo/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routers.Index(r)
	r.Run()
}
