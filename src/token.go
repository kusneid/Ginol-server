package src

import (
	"errors"
	"os/user"
	"sync"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GenerateToken() string {
	return uuid.New().String()
}

var (
	tokenStore = make(map[string]string) //[token]username
	tokenMutex = sync.RWMutex{}
)

func StoreToken(token, username string) {
	tokenMutex.Lock()
	defer tokenMutex.Unlock()
	tokenStore[token] = username
}

func GetUsernameByToken(token string) (string, bool) {
	tokenMutex.RLock()
	defer tokenMutex.RUnlock()
	username, exists := tokenStore[token]
	return username, exists
}

func UsernameExists(substring string) (bool, error) {
	var usr user.User

	// Используем ILIKE для нечувствительного к регистру частичного совпадения в PostgreSQL
	result := db.Where("username = ?", substring).First(&usr)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}

	return true, nil
}
