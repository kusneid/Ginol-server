package src

import (
	"sync"

	"github.com/google/uuid"
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
