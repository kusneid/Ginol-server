package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()

	r.POST("/api/loginServerHandler", func(c *gin.Context) {
		type Credentials struct {
			gorm.Model
			Username string `gorm:"username" json:"username"`
			Password string `gorm:"password" json:"password"`
		}
		var creds Credentials
		if err := c.ShouldBindJSON(&creds); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}
		fmt.Print(creds.Username, creds.Password)
	})

	r.Run(":8081")
}
