package src

import (
	"errors"
	"log"
	"os"

	"github.com/kusneid/Ginol/backend/user"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func VerifyPassword(c *user.Credentials, providedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(c.Password), []byte(providedPassword))
}

var db *gorm.DB

func InitDB() {
	dsn := os.Getenv("DATABASE_DATA")
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("failed to connect database", err)
	}
}

func NicknameExists(nickname string) (bool, error) {
	var user user.User
	result := db.Where("nickname = ?", nickname).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}

func LoginCheck(c *user.Credentials) (bool, error) {
	var creds user.Credentials

	if err := db.Where("username = ?", c.Username).First(&creds).Error; err != nil {
		log.Println("can't find user")
		return false, err
	}

	if err := VerifyPassword(&creds, c.Password); err != nil {
		log.Println("password incorrect")
		return false, err
	}

	return true, nil

}
