package model

import (
	"gorm.io/gorm"
)

type UserItem struct {
	gorm.Model
	Name  string `msgp:"name" db:"name"`
	Email string `msgp:"email" db:"email"`
}
