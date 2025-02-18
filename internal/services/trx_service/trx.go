package trxservice

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/nurmuh-alhakim18/evermos-project/helpers"
	alamatinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/alamat_interface"
	detailtrxinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/detail_trx_interface"
	fotoprodukinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/foto_produk_interface"
	kategoriinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/kategori_interface"
	logprodukinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/log_produk_interface"
	produkinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/produk_interface"
	tokointerface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/toko_interface"
	trxinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/trx_interface"
	detailtrxmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/detail_trx_model"
	logprodukmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/log_produk_model"
	trxmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/trx_model"
)

type TrxService struct {
	TrxRepository     trxinterface.TrxRepostoryInterface
	AlamatService     alamatinterface.AlamatServiceInterface
	ProductService    produkinterface.ProdukServiceInterface
	LogProductService logprodukinterface.LogProdukServiceInterface
	DetailTrxService  detailtrxinterface.DetailTrxServiceInterface
	TokoService       tokointerface.TokoServiceInterface
	KategoriService   kategoriinterface.KategoriServiceInterface
	FotoProdukService fotoprodukinterface.FotoProdukServiceInterface
}

func (s *TrxService) CreateTrx(ctx context.Context, userID int, req trxmodel.TrxReq) (*int, error) {
	shortID, err := helpers.GenerateShortID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate ID: %v", err)
	}

	invoiceCode := fmt.Sprintf("INV-%s%d", shortID, time.Now().Unix())

	trx := trxmodel.Trx{
		IdUser:           userID,
		HargaTotal:       0,
		AlamatPengiriman: req.AlamatPengiriman,
		MethodBayar:      req.MethodBayar,
		KodeInvoice:      invoiceCode,
	}

	trxID, err := s.TrxRepository.CreateTrx(ctx, &trx)
	if err != nil {
		return nil, fmt.Errorf("failed to create trx: %v", err)
	}

	var hargaTotal int
	for _, detailTrx := range req.DetailTrx {
		produk, err := s.ProductService.GetProdukByID(ctx, detailTrx.ProductID)
		if err != nil {
			return nil, fmt.Errorf("failed to get produk: %v", err)
		}

		if detailTrx.Kuantitas > produk.Stok {
			return nil, fmt.Errorf("failed to process transaction as the stock is not ready")
		}

		err = s.ProductService.UpdateStokAfterTransaction(ctx, produk.ID, detailTrx.Kuantitas)
		if err != nil {
			return nil, fmt.Errorf("failed to update stok: %v", err)
		}

		harga, err := strconv.Atoi(produk.HargaKonsumen)
		if err != nil {
			return nil, fmt.Errorf("failed to convert harga: %v", err)
		}

		hargaTotal += harga * detailTrx.Kuantitas

		logProduk := logprodukmodel.LogProduk{
			IdProduk:      produk.ID,
			NamaProduk:    produk.NamaProduk,
			Slug:          produk.Slug,
			HargaReseller: produk.HargaReseller,
			HargaKonsumen: produk.HargaKonsumen,
			Stok:          produk.Stok,
			Deskripsi:     produk.Deskripsi,
			IdToko:        produk.IdToko,
			IdKategori:    produk.IdKategori,
		}

		logProdukID, err := s.LogProductService.CreateLogProduk(ctx, logProduk)
		if err != nil {
			return nil, fmt.Errorf("failed to create log produk: %v", err)
		}

		hargaTotalPerItem := harga * detailTrx.Kuantitas
		detailTrx := detailtrxmodel.DetailTrx{
			IdTrx:       *trxID,
			IdLogProduk: *logProdukID,
			IdToko:      produk.IdToko,
			Kuantitas:   detailTrx.Kuantitas,
			HargaTotal:  hargaTotalPerItem,
		}

		_, err = s.DetailTrxService.CreateDetailTrx(ctx, detailTrx)
		if err != nil {
			return nil, fmt.Errorf("failed to create detail trx: %v", err)
		}
	}

	err = s.TrxRepository.UpdateTotalPrice(ctx, *trxID, hargaTotal)
	if err != nil {
		return nil, fmt.Errorf("faield to update total price: %v", err)
	}

	return trxID, err
}

func (s *TrxService) GetTrxByUserID(ctx context.Context, userID, limit, page int, search string) ([]trxmodel.GetTrxResp, error) {
	offset := (page - 1) * limit

	trxs, err := s.TrxRepository.GetTrxByUserID(ctx, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get trxs: %v", err)
	}

	var resp []trxmodel.GetTrxResp
	var detailTrxsResp []trxmodel.DetailTrxResp
	for _, trx := range trxs {
		alamat, err := s.AlamatService.GetAlamatByID(ctx, trx.AlamatPengiriman)
		if err != nil {
			return nil, fmt.Errorf("failed to get alamat: %v", err)
		}

		a := trxmodel.AlamatKirim{
			ID:           alamat.ID,
			JudulAlamat:  alamat.JudulAlamat,
			NamaPenerima: alamat.NamaPenerima,
			NoTelp:       alamat.NoTelp,
			DetailAlamat: alamat.DetailAlamat,
		}

		detailTrxs, err := s.DetailTrxService.GetDetailTrxByTrxID(ctx, trx.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get detail trx: %v", err)
		}

		for _, detailTrx := range detailTrxs {
			logProduk, err := s.LogProductService.GetLogProdukByID(ctx, detailTrx.IdLogProduk)
			if err != nil {
				return nil, fmt.Errorf("failed to get produk: %v", err)
			}

			toko, err := s.TokoService.GetTokoByID(ctx, detailTrx.IdToko)
			if err != nil {
				return nil, fmt.Errorf("failed to get toko: %v", err)
			}

			t := trxmodel.Toko{
				ID:       toko.ID,
				NamaToko: toko.NamaToko,
				URLFoto:  toko.URLFoto,
			}

			kategori, err := s.KategoriService.GetKategoriByID(ctx, logProduk.IdKategori)
			if err != nil {
				return nil, fmt.Errorf("failed to get kategori: %v", err)
			}

			k := trxmodel.Kategori{
				ID:           kategori.ID,
				NamaKategori: kategori.NamaKategori,
			}

			fotoProduks, err := s.FotoProdukService.GetFotoProdukByProdukID(ctx, logProduk.IdProduk)
			if err != nil {
				return nil, fmt.Errorf("failed to get foto produk: %v", err)
			}

			var fp []trxmodel.FotoProduk
			for _, fotoProduk := range fotoProduks {
				foto := trxmodel.FotoProduk{
					ID:       fotoProduk.ID,
					IdProduk: fotoProduk.IdProduk,
					URL:      fotoProduk.URL,
				}

				fp = append(fp, foto)
			}

			p := trxmodel.Produk{
				ID:            logProduk.ID,
				NamaProduk:    logProduk.NamaProduk,
				Slug:          logProduk.Slug,
				HargaReseller: logProduk.HargaReseller,
				HargaKonsumen: logProduk.HargaKonsumen,
				Deskripsi:     logProduk.Deskripsi,
				Toko:          t,
				Kategori:      k,
				Photos:        fp,
			}

			dt := trxmodel.DetailTrxResp{
				Product:   p,
				Toko:      t,
				Kuantitas: detailTrx.Kuantitas,
				Harga:     detailTrx.HargaTotal,
			}

			detailTrxsResp = append(detailTrxsResp, dt)
		}

		t := trxmodel.GetTrxResp{
			ID:          trx.ID,
			IdUser:      trx.IdUser,
			HargaTotal:  trx.HargaTotal,
			KodeInvoice: trx.KodeInvoice,
			MethodBayar: trx.MethodBayar,
			AlamatKirim: a,
			DetailTrx:   detailTrxsResp,
		}

		resp = append(resp, t)
	}

	if search != "" {
		var searchResp []trxmodel.GetTrxResp
		for _, r := range resp {
			for _, dt := range r.DetailTrx {
				if strings.Contains(strings.ToLower(dt.Product.NamaProduk), strings.ToLower(search)) {
					searchResp = append(searchResp, r)
				}
			}
		}

		return searchResp, nil
	}

	return resp, nil
}

func (s *TrxService) GetTrxByID(ctx context.Context, trxID int) (*trxmodel.GetTrxResp, error) {
	trx, err := s.TrxRepository.GetTrxByID(ctx, trxID)
	if err != nil {
		return nil, fmt.Errorf("failed to get trx: %v", err)
	}

	alamat, err := s.AlamatService.GetAlamatByID(ctx, trx.AlamatPengiriman)
	if err != nil {
		return nil, fmt.Errorf("failed to get alamat: %v", err)
	}

	a := trxmodel.AlamatKirim{
		ID:           alamat.ID,
		JudulAlamat:  alamat.JudulAlamat,
		NamaPenerima: alamat.NamaPenerima,
		NoTelp:       alamat.NoTelp,
		DetailAlamat: alamat.DetailAlamat,
	}

	detailTrxs, err := s.DetailTrxService.GetDetailTrxByTrxID(ctx, trx.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get detail trx: %v", err)
	}

	var detailTrxsResp []trxmodel.DetailTrxResp
	for _, detailTrx := range detailTrxs {
		logProduk, err := s.LogProductService.GetLogProdukByID(ctx, detailTrx.IdLogProduk)
		if err != nil {
			return nil, fmt.Errorf("failed to get produk: %v", err)
		}

		toko, err := s.TokoService.GetTokoByID(ctx, detailTrx.IdToko)
		if err != nil {
			return nil, fmt.Errorf("failed to get toko: %v", err)
		}

		t := trxmodel.Toko{
			ID:       toko.ID,
			NamaToko: toko.NamaToko,
			URLFoto:  toko.URLFoto,
		}

		kategori, err := s.KategoriService.GetKategoriByID(ctx, logProduk.IdKategori)
		if err != nil {
			return nil, fmt.Errorf("failed to get kategori: %v", err)
		}

		k := trxmodel.Kategori{
			ID:           kategori.ID,
			NamaKategori: kategori.NamaKategori,
		}

		fotoProduks, err := s.FotoProdukService.GetFotoProdukByProdukID(ctx, logProduk.IdProduk)
		if err != nil {
			return nil, fmt.Errorf("failed to get foto produk: %v", err)
		}

		var fp []trxmodel.FotoProduk
		for _, fotoProduk := range fotoProduks {
			foto := trxmodel.FotoProduk{
				ID:       fotoProduk.ID,
				IdProduk: fotoProduk.IdProduk,
				URL:      fotoProduk.URL,
			}

			fp = append(fp, foto)
		}

		p := trxmodel.Produk{
			ID:            logProduk.ID,
			NamaProduk:    logProduk.NamaProduk,
			Slug:          logProduk.Slug,
			HargaReseller: logProduk.HargaReseller,
			HargaKonsumen: logProduk.HargaKonsumen,
			Deskripsi:     logProduk.Deskripsi,
			Toko:          t,
			Kategori:      k,
			Photos:        fp,
		}

		dt := trxmodel.DetailTrxResp{
			Product:   p,
			Toko:      t,
			Kuantitas: detailTrx.Kuantitas,
			Harga:     detailTrx.HargaTotal,
		}

		detailTrxsResp = append(detailTrxsResp, dt)
	}

	t := trxmodel.GetTrxResp{
		ID:          trx.ID,
		IdUser:      trx.IdUser,
		HargaTotal:  trx.HargaTotal,
		KodeInvoice: trx.KodeInvoice,
		MethodBayar: trx.MethodBayar,
		AlamatKirim: a,
		DetailTrx:   detailTrxsResp,
	}

	return &t, nil
}
