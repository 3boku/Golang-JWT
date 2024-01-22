package middleware

import (
	"github.com/3boku/Golang-JWT/database"
	"github.com/gin-gonic/gin"
)

func RedisMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		rdb := database.RedisInit()
		c.Set("rdb", rdb)
		c.Next()
	}
}
