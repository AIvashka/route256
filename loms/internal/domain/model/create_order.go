package model

import "errors"

var (
	ErrorMissingUserId      = errors.New("missing user ID")
	ErrorItemListEmpty      = errors.New("items list is empty")
	ErrorMissingSKU         = errors.New("missing SKU")
	ErrorOrderItemCountZero = errors.New("item count is zero")
	ErrorOrderID            = errors.New("orderID must be provided")
)

type CreateOrderRequest struct {
	ID         int64       `json:"user"`
	OrderItems []OrderItem `json:"items"`
}

type CreateOrderResponse struct {
	OrderID int64 `json:"orderID"`
}

func (orderRequest CreateOrderRequest) Validate() error {
	if orderRequest.ID == 0 {
		return ErrorMissingUserId
	}
	if len(orderRequest.OrderItems) == 0 {
		return ErrorItemListEmpty
	}
	for _, item := range orderRequest.OrderItems {
		if item.SKU == 0 {
			return ErrorMissingSKU
		}
		if item.Count == 0 {
			return ErrorOrderItemCountZero
		}
	}
	return nil
}
