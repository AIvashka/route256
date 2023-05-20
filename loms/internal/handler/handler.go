package handler

import (
	"gitlab.ozon.dev/alexeyivashka/homework/loms/internal/usecase"
	"net/http"
)

type Handler struct {
	createOrderUseCase usecase.CreateOrderUseCase
	listOrderUseCase   usecase.ListOrderUseCase
	orderPaidUseCase   usecase.OrderPaidUseCase
	cancelOrderUseCase usecase.CancelOrderUseCase
	stocksUseCase      usecase.StocksUseCase
}

// NewHandler creates a new Handler with the provided use cases.
func NewHandler(
	createOrderUseCase usecase.CreateOrderUseCase,
	listOrderUseCase usecase.ListOrderUseCase,
	orderPaidUseCase usecase.OrderPaidUseCase,
	cancelOrderUseCase usecase.CancelOrderUseCase,
	stocksUseCase usecase.StocksUseCase,
) *Handler {
	return &Handler{
		createOrderUseCase: createOrderUseCase,
		listOrderUseCase:   listOrderUseCase,
		orderPaidUseCase:   orderPaidUseCase,
		cancelOrderUseCase: cancelOrderUseCase,
		stocksUseCase:      stocksUseCase,
	}
}

func SetupRoutes(handler *Handler) {
	http.HandleFunc("/createOrder", handler.CreateOrder)
	http.HandleFunc("/listOrder", handler.ListOrder)
	http.HandleFunc("/orderPaid", handler.OrderPaid)
	http.HandleFunc("/cancelOrder", handler.CancelOrder)
	http.HandleFunc("/stocks", handler.Stocks)
}
