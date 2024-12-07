package database

import (
	"Nikcase/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectToDB(config *models.Configs) (*gorm.DB, error) {
	dbUri := "host=" + config.Db.Host + " port=" + config.Db.Port +
		" user=" + config.Db.User + " password=" + config.Db.Password + " dbname=" + config.Db.Database
	db, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return db, nil
}
