package api

import (
	"SecretHitlerBackend/environment"
	"SecretHitlerBackend/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAvailableRooms(config *environment.AppConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		rooms, err := model.GetAvailableRooms(config.DB)
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, rooms)
	}
}

func CreateRoom(config *environment.AppConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := model.GetUserFromContext(c, config.DB)
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		code, err := model.CreateRoom(user, config.DB)
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, code)
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

		c.JSON(http.StatusOK, gin.H{})
	}
}
