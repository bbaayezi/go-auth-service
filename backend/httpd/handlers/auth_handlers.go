package handlers

import (
	"crypto/md5"
	"fmt"
	"go-auth-service/auth"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
)

var (
	password = "abc123"
)

type Login struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func TestPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// bind posted data
		var login Login
		c.BindJSON(&login)
		c.JSON(http.StatusOK, gin.H{
			"message": "get posted data",
			"data":    login,
		})
	}
}

func LoginHandler(redisClient *redis.Client) gin.HandlerFunc {

	return func(c *gin.Context) {
		// bind posted data
		var login Login
		c.BindJSON(&login)

		// check the request header
		// if there is no challenge reponse, then generate a challenge
		if cRes := c.GetHeader("X-Challenge-Response"); cRes != "" {
			// TODO: compare the challenge response with server generated answer
			key := login.User + ".challenge"
			val, err := redisClient.Get(key).Result()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("key: %v, value: %v\n", key, val)
				// create challenge object
				obj := &auth.Challenge{
					ChallengeString: val,
					Hash:            md5.New(),
				}
				if obj.Validate(password, cRes) == true {
					c.JSON(http.StatusOK, gin.H{
						"status":  "ok",
						"message": "authentication success",
					})
					// update challenge
					newChallenge := auth.NewChallenge(8)
					// update redis db
					err := redisClient.Set(key, newChallenge.String(), time.Minute*5).Err()
					if err != nil {
						fmt.Println(err)
					}
				} else {
					c.JSON(http.StatusUnauthorized, gin.H{
						"status": "unauthorized",
					})
				}
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "internal service error",
			})
		} else {
			// generate a challenge
			challenge := &auth.Challenge{}
			err := redisClient.Set(login.User+".challenge", challenge.String(), 0).Err()
			if err != nil {
				fmt.Println(err)
			}
			c.JSON(http.StatusNetworkAuthenticationRequired, gin.H{
				"status": "need authentication",
				"data": gin.H{
					"challenge": challenge.String(),
				},
			})
		}
	}
}
