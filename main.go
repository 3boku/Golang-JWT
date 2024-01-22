package main

import (
	"github.com/3boku/Golang-JWT/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.RedisMiddleware())
}
