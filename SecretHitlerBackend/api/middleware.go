package api

import (
	"SecretHitlerBackend/environment"
	"SecretHitlerBackend/model"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware(config *environment.AppConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionKey, err := c.Cookie("session_key")
		if err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("no session key"))
			return
		}

		session, err := model.GetSession(sessionKey, config.DB)
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, errors.New("no session key"))
			return
		}

		c.Set("user_id", session.UserID)

		c.Next()
	}
}
