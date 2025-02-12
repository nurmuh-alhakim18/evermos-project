package healthservice

import (
	healthinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/health_interface"
)

type HealthService struct {
	Repository healthinterface.HealthRepositoryInterface
}

func (s *HealthService) HealthCheck() (string, error) {
	return "Service OK", nil
}
