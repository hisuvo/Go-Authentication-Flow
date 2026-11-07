package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(env *Config) *gorm.DB {
	dsn := env.Dns
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return nil
	}

	Migrate(db)
	fmt.Println("Successfully database connected done!")
	
	return db
}