package postgres

import (
	"github.com/Gurpreet-Bacancy/bcl/model"
)

type UserManager struct {
	app *Application
}

func (am UserManager) GetUserByEmail(email string) (model.UserItem, error) {
	var userItems model.UserItem

	return userItems, nil
}

func (am UserManager) GetUserByID(userID uint) (model.UserItem, error) {
	var userItems model.UserItem

	return userItems, nil
}
