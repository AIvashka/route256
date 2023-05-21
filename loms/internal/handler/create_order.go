package handler

import (
	"context"
	"net/http"

	"gitlab.ozon.dev/alexeyivashka/homework/libs/srvwrapper"
	"gitlab.ozon.dev/alexeyivashka/homework/loms/internal/domain/model"
)

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	wrapper := srvwrapper.NewWrapper[model.CreateOrderRequest, model.CreateOrderResponse](h.handleCreateOrderRequest)
	wrapper.ServeHTTP(w, r)
}

func (h *Handler) handleCreateOrderRequest(ctx context.Context, request model.CreateOrderRequest) (model.CreateOrderResponse, error) {
	orderID, err := h.createOrderUseCase.Execute(ctx, request.ID, request.OrderItems)
	if err != nil {
		return model.CreateOrderResponse{}, err
	}
	return model.CreateOrderResponse{OrderID: orderID}, nil
}
