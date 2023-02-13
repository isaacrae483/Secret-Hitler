package api

import (
	"SecretHitlerBackend/environment"
	"SecretHitlerBackend/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRooms(config *environment.AppConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "GET"})
	}
}

func CreateRoom(config *environment.AppConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		code, err := model.CreateRoom(config.DB)
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
		}
		c.JSON(http.StatusOK, gin.H{"code": code})
	}
}

func JoinRoom(config *environment.AppConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "JOIN"})
	}
}
