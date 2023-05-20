package storage

import (
	"context"

	"gitlab.ozon.dev/alexeyivashka/homework/loms/internal/domain/model"
	"go.uber.org/zap"
)

type WarehouseStorage struct {
	logger *zap.Logger
}

func NewMockWarehouseRepository(logger *zap.Logger) *WarehouseStorage {
	return &WarehouseStorage{
		logger: logger,
	}
}

func (w WarehouseStorage) FindAll(ctx context.Context) ([]*model.Warehouse, error) {
	return []*model.Warehouse{
		{
			ID: 1,
			Items: []model.WarehouseItem{
				{
					ItemID: 1001,
					Count:  50,
				},
				{
					ItemID: 1002,
					Count:  70,
				},
				{
					ItemID: 1003,
					Count:  100,
				},
			},
		},
		{
			ID: 2,
			Items: []model.WarehouseItem{
				{
					ItemID: 2001,
					Count:  30,
				},
				{
					ItemID: 2002,
					Count:  80,
				},
				{
					ItemID: 2003,
					Count:  150,
				},
			},
		},
		{
			ID: 3,
			Items: []model.WarehouseItem{
				{
					ItemID: 3001,
					Count:  20,
				},
				{
					ItemID: 3002,
					Count:  60,
				},
				{
					ItemID: 3003,
					Count:  200,
				},
			},
		},
	}, nil
}

func (w WarehouseStorage) Update(ctx context.Context, warehouse *model.Warehouse) error {
	return nil
}

func (w WarehouseStorage) CancelReservation(ctx context.Context, sku uint32, count uint16) error {
	return nil
}

func (w WarehouseStorage) Reserve(ctx context.Context, sku uint32, count uint16) error {
	return nil
}

func (w WarehouseStorage) Purchase(ctx context.Context, sku uint32, count uint16) error {
	return nil
}

func (w WarehouseStorage) GetReservedCount(ctx context.Context, sku uint32) (uint64, error) {
	return 2312341, nil
}
