package storage

import (
	"context"

	"gitlab.ozon.dev/alexeyivashka/homework/loms/internal/domain/model"
	"go.uber.org/zap"
)

type ItemStorage struct {
	logger *zap.Logger
}

func NewMockItemRepository(logger *zap.Logger) *ItemStorage {
	return &ItemStorage{
		logger: logger,
	}
}

func (i ItemStorage) FindBySKU(ctx context.Context, sku uint32) (*model.Item, error) {
	return &model.Item{
		SKU:   12345453,
		Stock: 12341431,
	}, nil
}

func (i ItemStorage) Update(ctx context.Context, item *model.Item) error {
	return nil
}

func (i ItemStorage) CheckAvailability(ctx context.Context, sku uint32, count uint16) (bool, error) {
	return true, nil
}
