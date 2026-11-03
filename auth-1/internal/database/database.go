package database

import (
	"authentication-test/internal/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cnfg *config.Config) (*gorm.DB, error) {

	dsn := cnfg.DatabaseUrl

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	
    if err != nil {
        return nil, err
    }

	Migrate(db)
	
    log.Println("Database connected successfully")

    return db, nil
}