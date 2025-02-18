package trxservice

import (
	"context"
	"fmt"
	"time"

	"github.com/nurmuh-alhakim18/evermos-project/helpers"
	trxinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/trx_interface"
	trxmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/trx_model"
)

type TrxService struct {
	TrxRepository trxinterface.TrxRepostoryInterface
}

func (s *TrxService) CreateTrx(ctx context.Context, userID int, req trxmodel.TrxReq) (*int, error) {
	shortID, err := helpers.GenerateShortID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate ID: %v", err)
	}

	invoiceCode := fmt.Sprintf("INV-%s%d", shortID, time.Now().Unix())

	trx := trxmodel.Trx{
		IdUser:           userID,
		HargaTotal:       0, // unimplented
		AlamatPengiriman: req.AlamatPengiriman,
		MethodBayar:      req.MethodBayar,
		KodeInvoice:      invoiceCode,
	}

	trxID, err := s.TrxRepository.CreateTrx(ctx, &trx)
	if err != nil {
		return nil, fmt.Errorf("failed to create trx: %v", err)
	}

	return trxID, err
}

func (s *TrxService) GetTrxByUserID(ctx context.Context, userID, limit, page int, search string) ([]trxmodel.Trx, error) {
	offset := (page - 1) * limit

	trxs, err := s.TrxRepository.GetTrxByUserID(ctx, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get trxs: %v", err)
	}

	return trxs, nil
}

func (s *TrxService) GetTrxByID(ctx context.Context, trxID int) (*trxmodel.Trx, error) {
	trx, err := s.TrxRepository.GetTrxByID(ctx, trxID)
	if err != nil {
		return nil, fmt.Errorf("failed to get trx: %v", err)
	}

	return trx, nil
}
