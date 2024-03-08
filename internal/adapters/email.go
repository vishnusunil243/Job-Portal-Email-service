package adapters

import "gorm.io/gorm"

type EmailAdapter struct {
	DB *gorm.DB
}

func NewEmailAdapter(db *gorm.DB) *EmailAdapter {
	return &EmailAdapter{
		DB: db,
	}
}
