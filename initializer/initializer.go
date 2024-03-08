package initializer

import (
	"github.com/vishnusunil243/Job-Portal-Email-service/internal/adapters"
	"github.com/vishnusunil243/Job-Portal-Email-service/internal/service"
	"gorm.io/gorm"
)

func Initializer(db *gorm.DB) *service.EmailService {
	adapter := adapters.NewEmailAdapter(db)
	service := service.NewEmailService(adapter)
	return service
}
