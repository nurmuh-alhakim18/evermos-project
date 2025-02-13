package alamatservice

import (
	"context"
	"fmt"

	alamatinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/alamat_interface"
	alamatmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/alamat_model"
)

type AlamatService struct {
	AlamatRepository alamatinterface.AlamatRepositoryInterface
}

func (s *AlamatService) CreateAlamat(ctx context.Context, req alamatmodel.Alamat) error {
	err := s.AlamatRepository.CreateAlamat(ctx, &req)
	if err != nil {
		return fmt.Errorf("failed to create alamat: %v", err)
	}

	return nil
}

func (s *AlamatService) GetAlamat(ctx context.Context, userID int) ([]alamatmodel.Alamat, error) {
	alamats, err := s.AlamatRepository.GetAlamat(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get alamats: %v", err)
	}

	return alamats, nil
}

func (s *AlamatService) GetAlamatByID(ctx context.Context, alamatID int) (*alamatmodel.Alamat, error) {
	alamat, err := s.AlamatRepository.GetAlamatByID(ctx, alamatID)
	if err != nil {
		return nil, fmt.Errorf("failed to get alamats: %v", err)
	}

	if alamat == nil {
		return nil, fmt.Errorf("alamat not exists: %v", err)
	}

	return alamat, nil
}

func (s *AlamatService) UpdateAlamat(ctx context.Context, alamatID int, req alamatmodel.UpdateAlamat) (*alamatmodel.UpdateAlamat, error) {
	alamat, err := s.AlamatRepository.UpdateAlamat(ctx, alamatID, req)
	if err != nil {
		return nil, fmt.Errorf("failed to update alamat: %v", err)
	}

	if alamat == nil {
		return nil, fmt.Errorf("alamat not exists: %v", err)
	}

	return alamat, nil
}

func (s *AlamatService) DeleteAlamat(ctx context.Context, alamatID int) error {
	err := s.AlamatRepository.DeleteAlamat(ctx, alamatID)
	if err != nil {
		return fmt.Errorf("failed to delete alamat: %v", err)
	}

	return nil
}
