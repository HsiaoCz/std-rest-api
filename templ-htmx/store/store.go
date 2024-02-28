package store

import (
	"database/sql"

	"github.com/HsiaoCz/std-rest-api/templ-htmx/types"
)

type Store interface {
	CreateCar(car *types.Car) (*types.Car, error)
	GetCars() ([]types.Car, error)
	DeleteCar(id string) error
	FindCarsByNameMakeOrBrand(search string) ([]types.Car, error)
}

type Storage struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}
