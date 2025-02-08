package services

import "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces"

type HealthService struct {
	Repository interfaces.HealthRepositoryInterface
}

func (s *HealthService) HealthCheck() (string, error) {
	return "Service OK", nil
}
