package usecase

import (
	"context"
	"errors"

	"gitlab.ozon.dev/alexeyivashka/homework/loms/internal/domain/model"
	"gitlab.ozon.dev/alexeyivashka/homework/loms/internal/domain/repository"
	"go.uber.org/zap"
)

type CreateOrderUseCase interface {
	Execute(ctx context.Context, userID int64, items []model.OrderItem) (int64, error)
}

type createOrderUseCase struct {
	orderRepo     repository.OrderRepository
	itemRepo      repository.ItemRepository
	warehouseRepo repository.WarehouseRepository
	logger        *zap.Logger
}

func NewCreateOrderUseCase(or repository.OrderRepository,
	ir repository.ItemRepository,
	wr repository.WarehouseRepository,
	logger *zap.Logger) CreateOrderUseCase {
	return &createOrderUseCase{
		orderRepo:     or,
		itemRepo:      ir,
		warehouseRepo: wr,
		logger:        logger,
	}
}

func (uc *createOrderUseCase) Execute(ctx context.Context, userID int64, items []model.OrderItem) (int64, error) {
	for _, item := range items {
		available, err := uc.itemRepo.CheckAvailability(ctx, item.SKU, item.Count)
		if err != nil {
			return 0, err
		}
		if !available {
			return 0, errors.New("item not available in requested quantity")
		}
	}

	order := &model.Order{
		UserID: userID,
		Items:  items,
		Status: model.New,
	}
	orderID, err := uc.orderRepo.Create(ctx, order)
	if err != nil {
		return 0, err
	}

	for _, item := range items {
		err = uc.warehouseRepo.Reserve(ctx, item.SKU, item.Count)
		if err != nil {
			return 0, err
		}
	}

	return orderID, nil
}
