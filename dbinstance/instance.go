package dbinstance

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dir string = "host=localhost user=Grove password=root dbname=farmacia port=5432"

var DB *gorm.DB

func CrearConexion() (*gorm.DB, error) {

	DB, err := gorm.Open(postgres.Open(dir), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return DB, nil
}
