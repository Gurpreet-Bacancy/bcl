package postgres

import (
	"github.com/Gurpreet-Bacancy/bcl/dbconn"
)

type Application struct {
	db *dbconn.Postgres
}

type Models struct {
	User        UserManager
	Coordinates CoordinateManager
}

func NewModels(db *dbconn.Postgres) *Models {
	app := &Application{db: db}
	return &Models{
		User:        UserManager{app},
		Coordinates: CoordinateManager{app},
	}
}
