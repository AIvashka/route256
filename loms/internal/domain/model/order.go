package model

import "time"

type OrderStatus string

const (
	New             OrderStatus = "new"
	AwaitingPayment OrderStatus = "awaiting payment"
	Failed          OrderStatus = "failed"
	Paid            OrderStatus = "paid"
	Cancelled       OrderStatus = "cancelled"
)

type Order struct {
	ID        int64
	UserID    int64
	Status    OrderStatus
	Items     []OrderItem
	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrderItem struct {
	SKU   uint32
	Count uint16
}

type ListOrderRequest struct {
	ID int64 `json:"orderID"`
}

type ListOrderResponse struct {
	Status     OrderStatus `json:"status"`
	UserID     int64       `json:"user"`
	OrderItems []OrderItem `json:"items"`
}

func (l ListOrderRequest) Validate() error {
	return nil
}
