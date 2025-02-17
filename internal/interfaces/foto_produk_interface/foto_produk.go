package fotoprodukinterface

import (
	"context"
	"mime/multipart"

	fotoprodukmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/foto_produk_model"
)

type FotoProdukRepositoryInterface interface {
	CreateFotoProduk(ctx context.Context, req *fotoprodukmodel.FotoProduk) error
	GetFotoProdukByProdukID(ctx context.Context, produkID int) ([]fotoprodukmodel.FotoProduk, error)
	DeleteFotoProdukByProdukID(ctx context.Context, produkID int) error
}

type FotoProdukServiceInterface interface {
	CreateFotoProduk(ctx context.Context, produkID int, photo *multipart.FileHeader) error
	GetFotoProdukByProdukID(ctx context.Context, produkID int) ([]fotoprodukmodel.FotoProduk, error)
	DeleteFotoProdukByProdukID(ctx context.Context, produkID int) error
}
