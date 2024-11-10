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

	log.SetFlags(log.Ltime | log.Lshortfile)

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	src.InitDB()

	r.POST("/api/loginServerHandler", func(c *gin.Context) {

		var creds user.Credentials
		if err := c.ShouldBindJSON(&creds); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}
		fmt.Println(creds.Username, creds.Password)

		authenticated, err := src.LoginCheck(&creds)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"bool": "false"})
			return
		}

		if !authenticated {
			c.JSON(http.StatusUnauthorized, gin.H{"bool": "false"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"bool": "true"})
	})

	r.POST("/api/registrationServerHandler", func(c *gin.Context) {
		var creds user.Credentials
		if err := c.ShouldBindJSON(&creds); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}
		src.RegHandler(&creds)
		c.JSON(http.StatusOK, gin.H{"bool": "true"})
	})

	r.Run(":2737")
}
