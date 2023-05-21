package storage

import (
	"context"

	"gitlab.ozon.dev/alexeyivashka/homework/checkout/internal/domain/model"
)

type MockCartRepository struct {
	CartItems []model.CartItem
}

func NewMockCartRepository() *MockCartRepository {
	return &MockCartRepository{
		CartItems: []model.CartItem{
			{
				SKU:   123,
				Count: 2,
			},
			{
				SKU:   456,
				Count: 1,
			},
			{
				SKU:   789,
				Count: 3,
			},
		},
	}
}

func (m *MockCartRepository) AddToCart(ctx context.Context, request *model.AddToCartRequest) error {
	return nil
}

func (m *MockCartRepository) RemoveFromCart(ctx context.Context, userID int64, SKU uint32, count uint16) error {
	return nil
}

func (m *MockCartRepository) ListCartItems(ctx context.Context, user int64) ([]model.CartItem, error) {
	return m.CartItems, nil
}

func (m *MockCartRepository) ClearCart(ctx context.Context, user int64) error {
	return nil
}
