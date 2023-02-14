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
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": code})
	}
}

func JoinRoom(config *environment.AppConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		jri := model.JoinRoomInput{}
		if err := c.BindJSON(&jri); err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		if err := jri.Join(config.DB); err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, gin.H{"hello": "JOIN"})
	}
}
