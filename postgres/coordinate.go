package postgres

import (
	"github.com/Gurpreet-Bacancy/bcl/model"
	"github.com/jinzhu/gorm"
)

type CoordinateManager struct {
	app *Application
}

func (am CoordinateManager) GetUserLocation(id uint) (model.Coordinates, error) {
	var coordinates model.Coordinates
	tx := am.app.db.DB.Where("user_id=?", id).First(&coordinates)
	if tx.RecordNotFound() {
		return model.Coordinates{}, gorm.ErrRecordNotFound
	} else if tx.Error != nil {
		return model.Coordinates{}, tx.Error
	}
	return coordinates, nil
}

func (am CoordinateManager) AddUserLocation(id uint, Coordinate model.Coordinates) error {
	tx := am.app.db.DB.Create(&Coordinate)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (am CoordinateManager) UpdateUserLocation(id uint, Coordinate model.Coordinates) error {
	dbCoordinate, err := am.GetUserLocation(id)
	if err != nil {
		return err
	}

	dbCoordinate.Latitude = Coordinate.Latitude
	dbCoordinate.Longitude = Coordinate.Longitude
	dbCoordinate.Altitude = Coordinate.Altitude

	am.app.db.DB.Save(&dbCoordinate)

	return nil
}
