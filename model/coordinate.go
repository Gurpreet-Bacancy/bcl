package model

import (
	"gorm.io/gorm"
)

// Coordinates struct to describe coordinates object.
type Coordinates struct {
	gorm.Model
	UserID    uint    `msgp:"user_id" db:"user_id,string"`
	Latitude  float32 `msgp:"latitude" db:"latitude"`
	Longitude float32 `msgp:"longitude" db:"longitude"`
	Altitude  float32 `msgp:"altitude" db:"altitude"`
}

// UserCoordinateItem struct to describe userCoordinateItem object.
type UserCoordinateItem struct {
	gorm.Model
	UserID    uint    `msgp:"user_id" db:"user_id,string"`
	Latitude  float32 `msgp:"latitude" db:"latitude"`
	Longitude float32 `msgp:"longitude" db:"longitude"`
	Altitude  float32 `msgp:"altitude" db:"altitude"`
	Distance  float32 `msgp:"distance" db:"distance"`
}

type CoordinateStore interface {
	GetUserLocation(userID uint) (Coordinates, error)
	AddUserLocation(userID uint, Coordinate Coordinates) error
	UpdateUserLocation(Coordinate Coordinates) error
	GetNearestUsers(coordinate Coordinates) ([]UserCoordinateItem, error)
}
