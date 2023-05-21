package loms_client

//go:generate easyjson -all

type OrderItem struct {
	SKU   uint32
	Count uint16
}

type CreateOrderRequest struct {
	UserID int    `json:"user"`
	Items  []Item `json:"items"`
}

type Item struct {
	SKU   int `json:"sku"`
	Count int `json:"count"`
}

type OrderPaidRequest struct {
	OrderID int `json:"orderID"`
}

type CancelOrderRequest struct {
	OrderID int `json:"orderID"`
}

type StocksRequest struct {
	SKU int `json:"SKU"`
}

type CreateOrderResponse struct {
	OrderID int64 `json:"orderID"`
}

type OrderPaidResponse struct {
}

type CancelOrderResponse struct {
}

type Stock struct {
	WarehouseID int64  `json:"warehouseID"`
	Count       uint64 `json:"count"`
}

type StocksResponse struct {
	Stocks []Stock `json:"stocks"`
}

type ListOrderRequest struct {
	ID int64 `json:"orderID"`
}

type ListOrderResponse struct {
	Status     OrderStatus `json:"status"`
	UserID     int64       `json:"user"`
	OrderItems []OrderItem `json:"items"`
}

type OrderStatus string

const (
	New             OrderStatus = "new"
	AwaitingPayment OrderStatus = "awaiting payment"
	Failed          OrderStatus = "failed"
	Paid            OrderStatus = "paid"
	Cancelled       OrderStatus = "cancelled"
)
