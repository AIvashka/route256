package model

import "context"

type CheckoutService interface {
	AddToCart(ctx context.Context, request *AddToCartRequest) (*AddToCartResponse, error)
	DeleteFromCart(ctx context.Context, request *DeleteFromCartRequest) (*DeleteFromCartResponse, error)
	ListCart(ctx context.Context, request *ListCartRequest) (*ListCartResponse, error)
	Purchase(ctx context.Context, request *PurchaseRequest) (*PurchaseResponse, error)
}

type AddToCartRequest struct {
	User  int64  `json:"user"`
	SKU   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

func (req AddToCartRequest) Validate() error {
	return nil
}

type AddToCartResponse struct {
}

type DeleteFromCartRequest struct {
	User  int64  `json:"user"`
	SKU   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

func (d DeleteFromCartRequest) Validate() error {
	return nil
}

type DeleteFromCartResponse struct {
}

type ListCartRequest struct {
	User int64 `json:"user"`
}

func (l ListCartRequest) Validate() error {
	return nil
}

type CartItem struct {
	SKU   uint32 `json:"sku"`
	Count uint16 `json:"count"`
}

type ListCartResponse struct {
	Items      []Item `json:"items"`
	TotalPrice uint16 `json:"totalPrice"`
}

type Item struct {
	SKU   uint32 `json:"sku"`
	Count uint16 `json:"count"`
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}

type PurchaseRequest struct {
	User int64 `json:"user"`
}

func (p PurchaseRequest) Validate() error {
	return nil
}

type PurchaseResponse struct {
	OrderID int64 `json:"orderID"`
}
