package main

import (
	"fmt"
	"go-auth-service/httpd/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
)

func main() {
	// connecting to redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping().Result()
	if err == nil {
		fmt.Println(pong)
		fmt.Println("Successfully connected to redis!")
	}
	defer client.Close()

	fmt.Println("Starting Authentication Service")
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to go authentication service")
	})

	// test
	r.POST("/test", handlers.TestPostHandler())

	// login
	// the user should solve the challenge at this time
	r.POST("/login", handlers.LoginHandler(client))
	// listening at port 8080
	r.Run()
}
