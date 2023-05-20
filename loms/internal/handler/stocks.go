package handler

import (
	"context"
	"net/http"

	"gitlab.ozon.dev/alexeyivashka/homework/loms/internal/domain/model"
	"gitlab.ozon.dev/alexeyivashka/homework/srvwrapper"
)

func (h *Handler) Stocks(w http.ResponseWriter, r *http.Request) {
	wrapper := srvwrapper.NewWrapper[model.StocksRequest, model.StocksResponse](h.handleStocksRequest)
	wrapper.ServeHTTP(w, r)
}

func (h *Handler) handleStocksRequest(ctx context.Context, request model.StocksRequest) (model.StocksResponse, error) {
	count, err := h.stocksUseCase.Execute(ctx, request.SKU)
	if err != nil {
		return model.StocksResponse{}, err
	}
	return model.StocksResponse{Stocks: []model.Stock{{WarehouseID: 1, Count: count}}}, nil
}
