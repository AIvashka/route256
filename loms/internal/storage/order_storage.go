package storage

import (
	"context"
	"time"

	"gitlab.ozon.dev/alexeyivashka/homework/loms/internal/domain/model"
	"go.uber.org/zap"
)

type OrderStorage struct {
	logger *zap.Logger
}

func NewMockOrderRepository(logger *zap.Logger) *OrderStorage {
	return &OrderStorage{
		logger: logger,
	}
}

func (o OrderStorage) Create(ctx context.Context, order *model.Order) (int64, error) {
	return 234124, nil
}

func (o OrderStorage) FindByID(ctx context.Context, orderID int64) (*model.Order, error) {
	return &model.Order{
		ID:     1234,
		UserID: 6543,
		Status: model.AwaitingPayment,
		Items: []model.OrderItem{
			{
				SKU:   1001,
				Count: 50,
			},
			{
				SKU:   1002,
				Count: 70,
			},
			{
				SKU:   1003,
				Count: 100,
			},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (o OrderStorage) Update(ctx context.Context, order *model.Order) error {
	return nil
}

func (o OrderStorage) Delete(ctx context.Context, orderID int64) error {
	return nil
}
