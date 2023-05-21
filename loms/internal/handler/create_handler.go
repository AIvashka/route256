package handler

import (
	"gitlab.ozon.dev/alexeyivashka/homework/libs/logger"
	"gitlab.ozon.dev/alexeyivashka/homework/loms/internal/storage"
	"gitlab.ozon.dev/alexeyivashka/homework/loms/internal/usecase"
)

func NewHandlerWithDependencies() *Handler {
	logger := logger.GetLogger()
	orderRepo := storage.NewMockOrderRepository(logger)
	itemRepo := storage.NewMockItemRepository(logger)
	warehouseRepo := storage.NewMockWarehouseRepository(logger)

	createOrderUC := usecase.NewCreateOrderUseCase(orderRepo, itemRepo, warehouseRepo, logger)
	listOrderUC := usecase.NewListOrderUseCase(orderRepo, logger)
	orderPaidUC := usecase.NewOrderPaidUseCase(orderRepo, warehouseRepo, logger)
	cancelOrderUC := usecase.NewCancelOrderUseCase(orderRepo, warehouseRepo, logger)
	stocksUC := usecase.NewStocksUseCase(itemRepo, warehouseRepo, logger)

	return NewHandler(createOrderUC, listOrderUC, orderPaidUC, cancelOrderUC, stocksUC)
}
