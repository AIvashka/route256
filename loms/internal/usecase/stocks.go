package usecase

import (
	"context"

	"gitlab.ozon.dev/alexeyivashka/homework/loms/internal/domain/repository"
	"go.uber.org/zap"
)

type StocksUseCase interface {
	Execute(ctx context.Context, sku uint32) (uint64, error)
}

type stocksUseCase struct {
	itemRepo      repository.ItemRepository
	warehouseRepo repository.WarehouseRepository
	logger        *zap.Logger
}

func NewStocksUseCase(ir repository.ItemRepository, wr repository.WarehouseRepository, logger *zap.Logger) StocksUseCase {
	return &stocksUseCase{
		itemRepo:      ir,
		warehouseRepo: wr,
		logger:        logger,
	}
}

func (uc *stocksUseCase) Execute(ctx context.Context, sku uint32) (uint64, error) {
	// TODO Update to handle multiple warehouses
	item, err := uc.itemRepo.FindBySKU(ctx, sku)
	if err != nil {
		return 0, err
	}
	reserved, err := uc.warehouseRepo.GetReservedCount(ctx, sku)
	if err != nil {
		return 0, err
	}
	available := item.Stock - reserved
	return available, nil
}
