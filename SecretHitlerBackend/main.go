package main

import (
	"SecretHitlerBackend/api"
	"SecretHitlerBackend/environment"
	"github.com/gin-gonic/gin"
)

func main() {
	config := environment.Connect(false)

	r := gin.Default()
	_ = r.SetTrustedProxies(nil)

	user := r.Group("/users")
	user.POST("/login", api.Login(config))
	user.POST("/signup", api.Signup(config))

	room := r.Group("/rooms")
	room.GET("/available", api.GetAvailableRooms(config))
	room.POST("/create", api.CreateRoom(config))
	room.PUT("/join", api.JoinRoom(config))

	if err := r.Run("localhost:8080"); err != nil {
		panic(err)
	}
}
