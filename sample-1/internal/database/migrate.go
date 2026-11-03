package database

import (
	blog "authentication-test/internal/domain/blogs"
	user "authentication-test/internal/domain/users"
	"fmt"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB)  {
	err := db.AutoMigrate(
		&user.User{},
		&blog.Blog{},
	);
	
	if err != nil {
		fmt.Println("Migrated error :", err)
	}
}