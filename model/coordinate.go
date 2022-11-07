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
