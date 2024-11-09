package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kusneid/Ginol-server/src"
	"github.com/kusneid/Ginol/backend/user"
)

func main() {
	r := gin.Default()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	r.POST("/api/loginServerHandler", func(c *gin.Context) {

		var creds user.Credentials
		if err := c.ShouldBindJSON(&creds); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}
		fmt.Print(creds.Username, creds.Password)
		if err := src.LoginCheck(&creds); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed", "detail": err.Error()})
			return
		}
	})

	r.Run(":2737")
}
