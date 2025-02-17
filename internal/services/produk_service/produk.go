package produkservice

import (
	"context"
	"fmt"
	"strings"
	"time"

	fotoprodukinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/foto_produk_interface"
	kategoriinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/kategori_interface"
	produkinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/produk_interface"
	tokointerface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/toko_interface"
	produkmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/produk_model"
)

type ProdukService struct {
	ProdukRepository  produkinterface.ProdukRepositoryInterface
	TokoService       tokointerface.TokoServiceInterface
	KategoriService   kategoriinterface.KategoriServiceInterface
	FotoProdukService fotoprodukinterface.FotoProdukServiceInterface
}

func (s *ProdukService) CreateProduk(ctx context.Context, userID int, req produkmodel.ProdukReq) (*int, error) {
	slugParts := strings.Split(req.NamaProduk, " ")
	slug := strings.Join(slugParts, "-")

	toko, err := s.TokoService.GetTokoByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get produk: %v", err)
	}

	tokoID := toko.ID
	newProduk := produkmodel.Produk{
		NamaProduk:    req.NamaProduk,
		Slug:          slug,
		HargaReseller: req.HargaReseller,
		HargaKonsumen: req.HargaKonsumen,
		Stok:          req.Stok,
		Deskripsi:     req.Deskripsi,
		IdToko:        tokoID,
		IdKategori:    req.IdKategori,
	}

	produkID, err := s.ProdukRepository.CreateProduk(ctx, &newProduk)
	if err != nil {
		return nil, fmt.Errorf("failed to create produk: %v", err)
	}

	if req.Photos != nil {
		for _, p := range req.Photos {
			err := s.FotoProdukService.CreateFotoProduk(ctx, *produkID, p)
			if err != nil {
				return nil, fmt.Errorf("failed to create foto produk: %v", err)
			}
		}
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

		var photos []produkmodel.FotoProduk
		fotoProduks, err := s.FotoProdukService.GetFotoProdukByProdukID(ctx, p.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get foto produk: %v", err)
		}

		for _, fp := range fotoProduks {
			p := produkmodel.FotoProduk{
				ID:       fp.ID,
				IdProduk: fp.IdProduk,
				URL:      fp.URL,
			}

			photos = append(photos, p)
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
			Photos: photos,
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

	var photos []produkmodel.FotoProduk
	fotoProduks, err := s.FotoProdukService.GetFotoProdukByProdukID(ctx, produkID)
	if err != nil {
		return nil, fmt.Errorf("failed to get foto produk: %v", err)
	}

	for _, fp := range fotoProduks {
		p := produkmodel.FotoProduk{
			ID:       fp.ID,
			IdProduk: fp.IdProduk,
			URL:      fp.URL,
		}

		photos = append(photos, p)
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
		Photos: photos,
	}, nil
}

func (s *ProdukService) UpdateProduk(ctx context.Context, produkID int, req produkmodel.ProdukReq) error {
	slugParts := strings.Split(req.NamaProduk, " ")
	slug := strings.Join(slugParts, "-")

	if req.Photos != nil {
		err := s.FotoProdukService.DeleteFotoProdukByProdukID(ctx, produkID)
		if err != nil {
			return fmt.Errorf("failed to delete foto produk: %v", err)
		}

		for _, p := range req.Photos {
			err = s.FotoProdukService.CreateFotoProduk(ctx, produkID, p)
			if err != nil {
				return fmt.Errorf("failed to create foto produk: %v", err)
			}
		}
	}

	produk := produkmodel.UpdateProduk{
		NamaProduk:    req.NamaProduk,
		Slug:          slug,
		HargaReseller: req.HargaReseller,
		HargaKonsumen: req.HargaKonsumen,
		Stok:          req.Stok,
		Deskripsi:     req.Deskripsi,
		IdKategori:    req.IdKategori,
		UpdatedAt:     time.Now(),
	}

	err := s.ProdukRepository.UpdateProduk(ctx, produkID, produk)
	if err != nil {
		return fmt.Errorf("failed to update produk: %v", err)
	}

	return nil
}

func (s *ProdukService) DeleteProduk(ctx context.Context, produkID int) error {
	err := s.FotoProdukService.DeleteFotoProdukByProdukID(ctx, produkID)
	if err != nil {
		return fmt.Errorf("failed to delete foto produk: %v", err)
	}

	err = s.ProdukRepository.DeleteProduk(ctx, produkID)
	if err != nil {
		return fmt.Errorf("failed to delete produk: %v", err)
	}

	return nil
}
