package model

import (
	"gorm.io/gorm"
)

type Coordinates struct {
	gorm.Model
	UserID    uint    `msgp:"user_id" db:"user_id,string"`
	Latitude  float32 `msgp:"latitude" db:"latitude"`
	Longitude float32 `msgp:"longitude" db:"longitude"`
	Altitude  float32 `msgp:"altitude" db:"altitude"`
}
type UserCoordinateItem struct {
	ID        uint    `msgp:"name" db:"userid"`
	Name      string  `msgp:"name" db:"name"`
	Email     string  `msgp:"email" db:"email"`
	Latitude  float32 `msgp:"latitude" db:"latitude"`
	Longitude float32 `msgp:"longitude" db:"longitude"`
	Altitude  float32 `msgp:"altitude" db:"altitude"`
}
type CoordinateStore interface {
	GetUserLocation(userID uint) (Coordinates, error)
	AddUserLocation(userID uint, Coordinate Coordinates) error
	UpdateUserLocation(Coordinate Coordinates) error
	GetNearestUsers(Coordinate Coordinates, num int) ([]UserCoordinateItem, error)
}
