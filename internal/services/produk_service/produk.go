package produkservice

import (
	"context"
	"fmt"
	"strings"
	"time"

	kategoriinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/kategori_interface"
	produkinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/produk_interface"
	tokointerface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/toko_interface"
	produkmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/produk_model"
)

type ProdukService struct {
	ProdukRepository produkinterface.ProdukRepositoryInterface
	TokoService      tokointerface.TokoServiceInterface
	KategoriService  kategoriinterface.KategoriServiceInterface
}

func (s *ProdukService) CreateProduk(ctx context.Context, userID int, req produkmodel.Produk) (*int, error) {
	slugParts := strings.Split(req.NamaProduk, " ")
	slug := strings.Join(slugParts, "-")

	toko, err := s.TokoService.GetTokoByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get produk: %v", err)
	}

	tokoID := toko.ID
	newProduk := req
	newProduk.Slug = slug
	newProduk.IdToko = tokoID

	produkID, err := s.ProdukRepository.CreateProduk(ctx, &newProduk)
	if err != nil {
		return nil, fmt.Errorf("failed to create produk: %v", err)
	}

	return produkID, err
}

func (s *ProdukService) GetProduks(ctx context.Context, queries produkmodel.GetProdukQueries) ([]produkmodel.GetProdukResp, error) {
	offset := (queries.Page - 1) * queries.Limit
	queriesDB := &produkmodel.GetProdukQueriesDB{
		NamaProduk: queries.NamaProduk,
		Limit:      queries.Limit,
		Offset:     offset,
		CategoryId: queries.CategoryId,
		TokoId:     queries.TokoId,
		MaxHarga:   queries.MaxHarga,
		MinHarga:   queries.MinHarga,
	}

	produks, err := s.ProdukRepository.GetProduks(ctx, *queriesDB)
	if err != nil {
		return nil, fmt.Errorf("failed to get produks: %v", err)
	}

	produkResp := []produkmodel.GetProdukResp{}
	for _, p := range produks {
		toko, err := s.TokoService.GetTokoByID(ctx, p.IdToko)
		if err != nil {
			return nil, fmt.Errorf("failed to get toko with id %d: %v", p.IdToko, err)
		}

		kategori, err := s.KategoriService.GetKategoriByID(ctx, p.IdKategori)
		if err != nil {
			return nil, fmt.Errorf("failed to get category with id %d: %v", p.IdKategori, err)
		}

		produkResp = append(produkResp, produkmodel.GetProdukResp{
			Produk: p,
			Toko: produkmodel.Toko{
				ID:       toko.ID,
				NamaToko: toko.NamaToko,
				URLFoto:  toko.URLFoto,
			},
			Kategori: produkmodel.Kategori{
				ID:           kategori.ID,
				NamaKategori: kategori.NamaKategori,
			},
		})
	}

	return produkResp, nil
}

func (s *ProdukService) GetProdukByID(ctx context.Context, produkID int) (*produkmodel.GetProdukResp, error) {
	produk, err := s.ProdukRepository.GetProdukByID(ctx, produkID)
	if err != nil {
		return nil, fmt.Errorf("failed to get produk: %v", err)
	}

	toko, err := s.TokoService.GetTokoByID(ctx, produk.IdToko)
	if err != nil {
		return nil, fmt.Errorf("failed to get toko: %v", err)
	}

	kategori, err := s.KategoriService.GetKategoriByID(ctx, produk.IdKategori)
	if err != nil {
		return nil, fmt.Errorf("failed to get category: %v", err)
	}

	return &produkmodel.GetProdukResp{
		Produk: *produk,
		Toko: produkmodel.Toko{
			ID:       toko.ID,
			NamaToko: toko.NamaToko,
			URLFoto:  toko.URLFoto,
		},
		Kategori: produkmodel.Kategori{
			ID:           kategori.ID,
			NamaKategori: kategori.NamaKategori,
		},
	}, nil
}

func (s *ProdukService) UpdateProduk(ctx context.Context, produkID int, req produkmodel.UpdateProduk) error {
	slugParts := strings.Split(req.NamaProduk, " ")
	slug := strings.Join(slugParts, "-")

	produk := req
	produk.Slug = slug
	produk.UpdatedAt = time.Now()

	err := s.ProdukRepository.UpdateProduk(ctx, produkID, produk)
	if err != nil {
		return fmt.Errorf("failed to update produk: %v", err)
	}

	return nil
}

func (s *ProdukService) DeleteProduk(ctx context.Context, produkID int) error {
	err := s.ProdukRepository.DeleteProduk(ctx, produkID)
	if err != nil {
		return fmt.Errorf("failed to delete produk: %v", err)
	}

	return nil
}
