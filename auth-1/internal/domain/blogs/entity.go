package blog

import "gorm.io/gorm"

type Blog struct {
    gorm.Model

    Title       string `gorm:"type:varchar(255);not null"`
    Description string `gorm:"type:text;not null"`
}