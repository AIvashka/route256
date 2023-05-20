package repository

import (
	"context"

	"gitlab.ozon.dev/alexeyivashka/homework/loms/internal/domain/model"
)

type WarehouseRepository interface {
	FindAll(ctx context.Context) ([]*model.Warehouse, error)
	Update(ctx context.Context, warehouse *model.Warehouse) error
	CancelReservation(ctx context.Context, sku uint32, count uint16) error
	Reserve(ctx context.Context, sku uint32, count uint16) error
	Purchase(ctx context.Context, sku uint32, count uint16) error
	GetReservedCount(ctx context.Context, sku uint32) (uint64, error)
}
