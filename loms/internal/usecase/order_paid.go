package usecase

import (
	"context"

	"gitlab.ozon.dev/alexeyivashka/homework/loms/internal/domain/model"
	repository2 "gitlab.ozon.dev/alexeyivashka/homework/loms/internal/domain/repository"
	"go.uber.org/zap"
)

type OrderPaidUseCase interface {
	Execute(ctx context.Context, orderID int64) error
}

type orderPaidUseCase struct {
	orderRepo     repository2.OrderRepository
	warehouseRepo repository2.WarehouseRepository
	logger        *zap.Logger
}

func NewOrderPaidUseCase(or repository2.OrderRepository, wr repository2.WarehouseRepository, logger *zap.Logger) OrderPaidUseCase {
	return &orderPaidUseCase{
		orderRepo:     or,
		warehouseRepo: wr,
		logger:        logger,
	}
}

func (uc *orderPaidUseCase) Execute(ctx context.Context, orderID int64) error {
	order, err := uc.orderRepo.FindByID(ctx, orderID)
	if err != nil {
		return err
	}

	order.Status = model.Paid
	err = uc.orderRepo.Update(ctx, order)
	if err != nil {
		return err
	}

	for _, item := range order.Items {
		err = uc.warehouseRepo.Purchase(ctx, item.SKU, item.Count)
		if err != nil {
			return err
		}
	}

	return nil
}
