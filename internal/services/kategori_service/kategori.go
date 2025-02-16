package kategoriservice

import (
	"context"
	"fmt"

	kategoriinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/kategori_interface"
	kategorimodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/kategori_model"
)

type KategoriService struct {
	KategoriRepository kategoriinterface.KategoriRepositoryInterface
}

func (s *KategoriService) CreateKategori(ctx context.Context, req kategorimodel.Kategori) (*int, error) {
	kategoriID, err := s.KategoriRepository.CreateKategori(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("failed to create kategori: %v", err)
	}

	return kategoriID, nil
}

func (s *KategoriService) GetKategoris(ctx context.Context) ([]kategorimodel.Kategori, error) {
	kategoris, err := s.KategoriRepository.GetKategoris(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get kategoris: %v", err)
	}

	return kategoris, nil
}

func (s *KategoriService) GetKategoriByID(ctx context.Context, kategoriID int) (*kategorimodel.Kategori, error) {
	kategori, err := s.KategoriRepository.GetKategoriByID(ctx, kategoriID)
	if err != nil {
		return nil, fmt.Errorf("failed to get kategori: %v", err)
	}

	return kategori, nil
}

func (s *KategoriService) UpdateKategori(ctx context.Context, kategoriID int, req kategorimodel.UpdateKategori) error {
	err := s.KategoriRepository.UpdateKategori(ctx, kategoriID, req)
	if err != nil {
		return fmt.Errorf("failed to update kategori: %v", err)
	}

	return nil
}

func (s *KategoriService) DeleteKategori(ctx context.Context, kategoriID int) error {
	err := s.KategoriRepository.DeleteKategori(ctx, kategoriID)
	if err != nil {
		return fmt.Errorf("failed to delete kategori: %v", err)
	}

	return nil
}
