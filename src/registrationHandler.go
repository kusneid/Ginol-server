package src

import "github.com/kusneid/Ginol/backend/user"

func RegHandler(u *user.Credentials) {
	db.Create(&u)
}
