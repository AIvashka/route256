package repository

import (
	"context"

	"gitlab.ozon.dev/alexeyivashka/homework/checkout/internal/domain/model"
)

type CartRepository interface {
	AddToCart(ctx context.Context, req *model.AddToCartRequest) error
	RemoveFromCart(ctx context.Context, userID int64, SKU uint32, count uint16) error
	ListCartItems(ctx context.Context, user int64) ([]model.CartItem, error)
	ClearCart(ctx context.Context, user int64) error
}
