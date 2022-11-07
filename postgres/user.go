package postgres

import (
	"github.com/Gurpreet-Bacancy/bcl/model"
	"github.com/jinzhu/gorm"
)

type UserManager struct {
	app *Application
}

func (am UserManager) GetUserByEmail(email string) (model.UserItem, error) {
	var userItems model.UserItem
	tx := am.app.db.DB.Where("email=?", email).First(&userItems)
	if tx.RecordNotFound() {
		return model.UserItem{}, gorm.ErrRecordNotFound
	} else if tx.Error != nil {
		return model.UserItem{}, tx.Error
	}
	return userItems, nil

}

func (am UserManager) GetUserByID(userID uint) (model.UserItem, error) {
	var userItems model.UserItem

	return userItems, nil
}
