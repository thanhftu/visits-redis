package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhftu/visits-redis/controller"
)

var (
	router = gin.Default()
)

func main() {
	router.GET("/", controller.GetVisits)
	router.Run(":8081")
}
