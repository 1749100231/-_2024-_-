package database

import (
	"JH_2024_MJJ/internal/model"
	"gorm.io/gorm"
)

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		model.User{},
		model.Article{},
		model.TokenTable{},
	)
}
