package usecase

import (
	"context"

	"gitlab.ozon.dev/alexeyivashka/homework/loms/internal/domain/model"
	"gitlab.ozon.dev/alexeyivashka/homework/loms/internal/domain/repository"
	"go.uber.org/zap"
)

type CancelOrderUseCase interface {
	Execute(ctx context.Context, orderID int64) error
}

type cancelOrderUseCase struct {
	orderRepo     repository.OrderRepository
	warehouseRepo repository.WarehouseRepository
	logger        *zap.Logger
}

func NewCancelOrderUseCase(or repository.OrderRepository, wr repository.WarehouseRepository, logger *zap.Logger) CancelOrderUseCase {
	return &cancelOrderUseCase{
		orderRepo:     or,
		warehouseRepo: wr,
		logger:        logger,
	}
}

func (uc *cancelOrderUseCase) Execute(ctx context.Context, orderID int64) error {
	order, err := uc.orderRepo.FindByID(ctx, orderID)
	if err != nil {
		return err
	}

	order.Status = model.Cancelled
	err = uc.orderRepo.Update(ctx, order)
	if err != nil {
		return err
	}

	for _, item := range order.Items {
		err = uc.warehouseRepo.CancelReservation(ctx, item.SKU, item.Count)
		if err != nil {
			return err
		}
	}

	return nil
}
