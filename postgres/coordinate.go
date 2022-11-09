package postgres

import (
	"fmt"

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

func (am CoordinateManager) GetNearestUsers(coordinate model.Coordinates) ([]model.UserCoordinateItem, error) {
	var userCoordinates []model.UserCoordinateItem
	query := fmt.Sprintf(`SELECT * , 3956 * 2 * ASIN(SQRT(
		POWER(SIN((%f - abs(dest.latitude)) * pi()/180 / 2),
		2) + COS(%f * pi()/180 ) * COS(abs(dest.latitude) *
		pi()/180) * POWER(SIN((%f - dest.longitude) *
		pi()/180 / 2), 2) )) as distance
		FROM coordinates dest 
		ORDER BY distance limit 10;`, coordinate.Latitude, coordinate.Latitude, coordinate.Longitude)
	rows, err := am.app.db.DB.Raw(query).Rows()
	if err != nil {
		fmt.Println("Error from get nearest db call", err)
	}
	defer rows.Close()
	for rows.Next() {
		var userCoordinate model.UserCoordinateItem
		// ScanRows scan a row into user
		am.app.db.DB.ScanRows(rows, &userCoordinate)

		// do something
		userCoordinates = append(userCoordinates, userCoordinate)
	}
	return userCoordinates, nil
}
