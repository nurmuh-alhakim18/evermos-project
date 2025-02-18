package tokoservice

import (
	"context"
	"fmt"
	"time"

	"github.com/nurmuh-alhakim18/evermos-project/helpers"
	tokointerface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/toko_interface"
	tokomodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/toko_model"
)

type TokoService struct {
	TokoRepository tokointerface.TokoRepositoryInterface
}

func (s *TokoService) CreateToko(ctx context.Context, req tokomodel.Toko) error {
	err := s.TokoRepository.CreateToko(ctx, &req)
	if err != nil {
		return fmt.Errorf("failed to create toko: %v", err)
	}

	return nil
}

func (s *TokoService) GetTokos(ctx context.Context, nama string, page, limit int) ([]tokomodel.Toko, error) {
	offset := (page - 1) * limit

	tokos, err := s.TokoRepository.GetTokos(ctx, nama, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get tokos: %v", err)
	}

	return tokos, nil
}

func (s *TokoService) GetTokoByUserID(ctx context.Context, userID int) (*tokomodel.Toko, error) {
	toko, err := s.TokoRepository.GetTokoByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get toko: %v", err)
	}

	return toko, nil
}

func (s *TokoService) GetTokoByID(ctx context.Context, tokoID int) (*tokomodel.Toko, error) {
	toko, err := s.TokoRepository.GetTokoByID(ctx, tokoID)
	if err != nil {
		return nil, fmt.Errorf("failed to get toko: %v", err)
	}

	return toko, nil
}

func (s *TokoService) UpdateToko(ctx context.Context, tokoID int, req tokomodel.UpdateTokoReq) error {
	var (
		url string
		err error
	)

	if req.Photo != nil {
		toko, err := s.TokoRepository.GetTokoByID(ctx, tokoID)
		if err != nil {
			return fmt.Errorf("failed to get toko: %v", err)
		}

		fotoURL := toko.URLFoto
		err = helpers.DeleteFromS3(fotoURL)
		if err != nil {
			return fmt.Errorf("failed to delete foto: %v", err)
		}

		url, err = helpers.UploadToS3(req.Photo)
		if err != nil {
			return err
		}
	}

	toko := tokomodel.UpdateToko{NamaToko: req.NamaToko, URLFoto: url, UpdatedAt: time.Now()}

	err = s.TokoRepository.UpdateToko(ctx, tokoID, toko)
	if err != nil {
		return fmt.Errorf("failed to update alamat: %v", err)
	}

	return nil
}
