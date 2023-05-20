package repository

import (
	"context"

	"gitlab.ozon.dev/alexeyivashka/homework/loms/internal/domain/model"
)

type OrderRepository interface {
	Create(ctx context.Context, order *model.Order) (int64, error)
	FindByID(ctx context.Context, orderID int64) (*model.Order, error)
	Update(ctx context.Context, order *model.Order) error
	Delete(ctx context.Context, orderID int64) error
}
