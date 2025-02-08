package healthS

import "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/healthI"

type HealthService struct {
	Repository healthI.HealthRepositoryInterface
}

func (s *HealthService) HealthCheck() (string, error) {
	return "Service OK", nil
}
