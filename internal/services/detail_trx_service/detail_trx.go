package detailtrxservice

import (
	"context"
	"fmt"

	detailtrxinterface "github.com/nurmuh-alhakim18/evermos-project/internal/interfaces/detail_trx_interface"
	detailtrxmodel "github.com/nurmuh-alhakim18/evermos-project/internal/models/detail_trx_model"
)

type DetailTrxService struct {
	DetailTrxRepository detailtrxinterface.DetailTrxRepositoryInterface
}

func (s *DetailTrxService) CreateDetailTrx(ctx context.Context, req detailtrxmodel.DetailTrx) (*int, error) {
	detailTrxID, err := s.DetailTrxRepository.CreateDetailTrx(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("failed to create detail trx: %v", err)
	}

	return detailTrxID, nil
}

func (s *DetailTrxService) GetDetailTrxByTrxID(ctx context.Context, trxID int) ([]detailtrxmodel.DetailTrx, error) {
	detailTrxs, err := s.DetailTrxRepository.GetDetailTrxByTrxID(ctx, trxID)
	if err != nil {
		return nil, fmt.Errorf("failed to get detail trx: %v", err)
	}

	return detailTrxs, nil
}
