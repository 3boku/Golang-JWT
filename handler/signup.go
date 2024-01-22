package handler

import (
	"context"
	"encoding/json"
	"github.com/3boku/Golang-JWT/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"net/http"
)

var (
	ctx = context.Background()
)

func SignUp(c *gin.Context) {
	rdb := c.MustGet("rdb").(*redis.Client)
	var user utils.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userJson, _ := json.Marshal(user)
	err := rdb.Set(ctx, user.ID, userJson, 0).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user to Redis"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "User saved to Redis"})
}
