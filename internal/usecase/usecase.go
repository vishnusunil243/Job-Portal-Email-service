package usecase

import "github.com/vishnusunil243/Job-Portal-Email-service/internal/adapters"

type CompanyUseCases struct {
	Adapter adapters.EmailInterface
}
