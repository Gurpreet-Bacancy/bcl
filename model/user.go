package model

import (
	"gorm.io/gorm"
)

type UserItem struct {
	gorm.Model
	Name  string `msgp:"name" db:"name"`
	Email string `msgp:"email" db:"email"`
}

type UserStore interface {
	GetUserByEmail(email string) (UserItem, error)
	GetUserByID(userID uint) (UserItem, error)
}
