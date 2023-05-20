package repository

import (
	"context"

	"gitlab.ozon.dev/alexeyivashka/homework/loms/internal/domain/model"
)

type ItemRepository interface {
	FindBySKU(ctx context.Context, sku uint32) (*model.Item, error)
	Update(ctx context.Context, item *model.Item) error
	CheckAvailability(ctx context.Context, sku uint32, count uint16) (bool, error)
}
