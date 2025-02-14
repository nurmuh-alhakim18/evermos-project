package wilayahservice

import (
	"context"
	"fmt"

	wilayahinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/wilayah_interface"
	usermodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/user_model"
)

type WilayahService struct {
	WilayahRepository wilayahinterface.WilayahRepositoryInterface
}

func (s *WilayahService) GetProvincies(ctx context.Context, search string, page, limit int) ([]usermodel.Provinsi, error) {
	offset := (page - 1) * limit

	provincies, err := s.WilayahRepository.GetProvinces(ctx, search, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get provincies: %v", err)
	}

	return provincies, nil
}

func (s *WilayahService) GetCities(ctx context.Context, provinceID string) ([]usermodel.Kota, error) {
	cities, err := s.WilayahRepository.GetCities(ctx, provinceID)
	if err != nil {
		return nil, fmt.Errorf("failed to get cities: %v", err)
	}

	return cities, nil
}

func (s *WilayahService) GetProvinceDetail(ctx context.Context, provinceID string) (*usermodel.Provinsi, error) {
	province, err := s.WilayahRepository.GetProvinceDetail(ctx, provinceID)
	if err != nil {
		return nil, fmt.Errorf("failed to get province: %v", err)
	}

	return province, nil
}

func (s *WilayahService) GetCityDetail(ctx context.Context, cityID string) (*usermodel.Kota, error) {
	city, err := s.WilayahRepository.GetCityDetail(ctx, cityID)
	if err != nil {
		return nil, fmt.Errorf("failed to get province: %v", err)
	}

	return city, nil
}
