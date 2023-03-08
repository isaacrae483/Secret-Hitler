package api

import (
	"SecretHitlerBackend/environment"
	"SecretHitlerBackend/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(config *environment.AppConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		credentials := model.Credentials{}
		if err := c.BindJSON(&credentials); err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		session, err := credentials.Login(config.DB)
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.SetCookie("session_key", session.Key, 3600, "/", "localhost", false, true)

		c.JSON(http.StatusOK, gin.H{})
	}
}

func Signup(config *environment.AppConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		credentials := model.Credentials{}
		if err := c.BindJSON(&credentials); err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		session, err := credentials.Signup(config.DB)
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.SetCookie("session_key", session.Key, 3600, "/", "localhost", false, true)

		c.JSON(http.StatusOK, gin.H{})
	}
}
