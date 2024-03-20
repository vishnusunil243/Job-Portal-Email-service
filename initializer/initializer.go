package initializer

import (
	"github.com/vishnusunil243/Job-Portal-Email-service/internal/adapters"
	"github.com/vishnusunil243/Job-Portal-Email-service/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func Initializer(db *mongo.Database) *service.EmailService {
	adapter := adapters.NewEmailAdapter(db)
	service := service.NewEmailService(adapter)
	return service
}
