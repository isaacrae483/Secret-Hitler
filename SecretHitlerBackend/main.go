package main

import (
	"SecretHitlerBackend/api"
	"SecretHitlerBackend/environment"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	config := environment.Connect(false)

	r := gin.Default()
	_ = r.SetTrustedProxies(nil)

	rAuth := r.Group("")
	rAuth.Use(api.AuthMiddleware(config))

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"test": "SUCCESS"})
	})

	user := r.Group("/users")
	user.POST("/login", api.Login(config))
	user.POST("/signup", api.Signup(config))

	room := rAuth.Group("/rooms")
	room.GET("/available", api.GetAvailableRooms(config))
	room.POST("/create", api.CreateRoom(config))
	room.PUT("/join", api.JoinRoom(config))

	if err := r.Run("localhost:8080"); err != nil {
		panic(err)
	}

}
