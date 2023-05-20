package usecase

import (
	"context"

	"gitlab.ozon.dev/alexeyivashka/homework/loms/internal/domain/model"
	"gitlab.ozon.dev/alexeyivashka/homework/loms/internal/domain/repository"
	"go.uber.org/zap"
)

type ListOrderUseCase interface {
	Execute(ctx context.Context, orderID int64) (*model.Order, error)
}

type listOrderUseCase struct {
	orderRepo repository.OrderRepository
	logger    *zap.Logger
}

func NewListOrderUseCase(or repository.OrderRepository, logger *zap.Logger) ListOrderUseCase {
	return &listOrderUseCase{
		orderRepo: or,
		logger:    logger,
	}
}

func (uc *listOrderUseCase) Execute(ctx context.Context, orderID int64) (*model.Order, error) {
	order, err := uc.orderRepo.FindByID(ctx, orderID)
	if err != nil {
		return nil, err
	}
	return order, nil
}
