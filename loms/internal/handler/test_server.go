package handler

import (
	"net/http"

	"gitlab.ozon.dev/alexeyivashka/homework/loms/internal/storage"
	"gitlab.ozon.dev/alexeyivashka/homework/loms/internal/usecase"
	"go.uber.org/zap"
)

func StartTestServer() *http.Server {
	logger := zap.NewExample() // или используйте ваш собственный логгер
	orderRepo := storage.NewMockOrderRepository(logger)
	itemRepo := storage.NewMockItemRepository(logger)
	warehouseRepo := storage.NewMockWarehouseRepository(logger)

	createOrderUC := usecase.NewCreateOrderUseCase(orderRepo, itemRepo, warehouseRepo, logger)
	listOrderUC := usecase.NewListOrderUseCase(orderRepo, logger)
	orderPaidUC := usecase.NewOrderPaidUseCase(orderRepo, warehouseRepo, logger)
	cancelOrderUC := usecase.NewCancelOrderUseCase(orderRepo, warehouseRepo, logger)
	stocksUC := usecase.NewStocksUseCase(itemRepo, warehouseRepo, logger)

	handler := NewHandler(createOrderUC, listOrderUC, orderPaidUC, cancelOrderUC, stocksUC)

	http.HandleFunc("/createOrder", handler.CreateOrder)
	http.HandleFunc("/listOrder", handler.ListOrder)
	http.HandleFunc("/orderPaid", handler.OrderPaid)
	http.HandleFunc("/cancelOrder", handler.CancelOrder)
	http.HandleFunc("/stocks", handler.Stocks)

	server := &http.Server{
		Addr: ":8080",
	}

	go func() {
		logger.Info("Starting Test Server")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("listen:%+s\n", zap.Error(err))
		}
	}()

	return server
}
