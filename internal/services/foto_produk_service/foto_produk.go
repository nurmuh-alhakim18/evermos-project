package fotoprodukservice

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/nurmuh-alhakim18/evermos-project/helpers"
	fotoprodukinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/foto_produk_interface"
	fotoprodukmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/foto_produk_model"
)

type FotoProdukService struct {
	FotoProdukRepository fotoprodukinterface.FotoProdukRepositoryInterface
}

func (s *FotoProdukService) CreateFotoProduk(ctx context.Context, produkID int, photo *multipart.FileHeader) error {
	url, err := helpers.UploadToS3(photo)
	if err != nil {
		return fmt.Errorf("failed to upload photo: %v", err)
	}

	req := fotoprodukmodel.FotoProduk{
		IdProduk: produkID,
		URL:      url,
	}

	err = s.FotoProdukRepository.CreateFotoProduk(ctx, &req)
	if err != nil {
		return fmt.Errorf("failed to create foto produk: %v", err)
	}

	return nil
}

func (s *FotoProdukService) GetFotoProdukByProdukID(ctx context.Context, produkID int) ([]fotoprodukmodel.FotoProduk, error) {
	fotoProduks, err := s.FotoProdukRepository.GetFotoProdukByProdukID(ctx, produkID)
	if err != nil {
		return nil, fmt.Errorf("failed to get foto produks: %v", err)
	}

	return fotoProduks, nil
}

func (s *FotoProdukService) DeleteFotoProdukByProdukID(ctx context.Context, produkID int) error {
	fotoProduks, err := s.GetFotoProdukByProdukID(ctx, produkID)
	if err != nil {
		return fmt.Errorf("failed to get foto produks: %v", err)
	}

	for _, fp := range fotoProduks {
		err = helpers.DeleteFromS3(fp.URL)
		if err != nil {
			return fmt.Errorf("failed to delete foto produk: %v", err)
		}
	}

	err = s.FotoProdukRepository.DeleteFotoProdukByProdukID(ctx, produkID)
	if err != nil {
		return fmt.Errorf("failed to delete foto produk: %v", err)
	}

	return nil
}
