package postgres

import (
	"github.com/Gurpreet-Bacancy/bcl/model"
)

type CoordinateManager struct {
	app *Application
}

func (am CoordinateManager) GetUserLocation(id uint) (model.Coordinates, error) {
	var coordinates model.Coordinates

	return coordinates, nil
}

func (am CoordinateManager) AddUserLocation(id uint, Coordinate model.Coordinates) error {
	return nil
}

func (am CoordinateManager) UpdateUserLocation(id uint, Coordinate model.Coordinates) error {
	return nil
}
