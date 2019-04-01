package db

import (
	"new-platform-dashboard/config"
	"new-platform-dashboard/db/entities"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
)

var DB []*gorm.DB

func Init() {
	for i := range config.Databases {
		var database = &config.Databases[i]

		db, err := gorm.Open(database.DriverName, database.ConnectionString)
		if err != nil {
			panic("failed to connect database")
		}

		db.DB().SetMaxOpenConns(database.MaxConnectionOpen)
		if config.App.Env == "dev" {
			db.LogMode(true)
		}

		log.WithFields(log.Fields{
			"config": database,
		}).Info("Connected to database")

		gormMigration(database.Name, db)
		//append database to array
		DB = append(DB, db)
	}
}

//register entity for created table
func gormMigration(dbName string, db *gorm.DB) {
	// if dbName == "gosample_db" {
	db.AutoMigrate(&entities.Example{}, &entities.User{})
	// }
}
