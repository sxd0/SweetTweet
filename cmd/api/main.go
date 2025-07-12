package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sxd0/SweetTweet/internal/api/grpcclient"
)

func main() {
	grpcclient.InitAuthClient()
	grpcclient.InitUserClient()

	r := gin.Default()

	r.POST("/register", func(c *gin.Context) {
		var req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.BindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "bad request"})
			return
		}
		res, err := grpcclient.Register(c, req.Email, req.Password)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, res)
	})

	r.POST("/login", func(c *gin.Context) {
		var req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.BindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "bad request"})
			return
		}
		res, err := grpcclient.Login(c, req.Email, req.Password)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, res)
	})

	r.GET("/me", func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(401, gin.H{"error": "unauthorized"})
			return
		}
		res, err := grpcclient.Me(c, token)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, res)
	})

	port := os.Getenv("GIN_PORT")
	if port == "" {
		port = ":8080"
	}
	log.Println("API Gateway running on", port)
	r.Run(port)
}