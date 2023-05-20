package handler

import (
	"context"
	"net/http"

	"gitlab.ozon.dev/alexeyivashka/homework/loms/internal/domain/model"
	"gitlab.ozon.dev/alexeyivashka/homework/srvwrapper"
)

func (h *Handler) CancelOrder(w http.ResponseWriter, r *http.Request) {
	wrapper := srvwrapper.NewWrapper[model.CancelOrderRequest, model.EmptyResponse](h.handleCancelOrderRequest)
	wrapper.ServeHTTP(w, r)
}

func (h *Handler) handleCancelOrderRequest(ctx context.Context, request model.CancelOrderRequest) (model.EmptyResponse, error) {
	err := h.cancelOrderUseCase.Execute(ctx, request.ID)
	if err != nil {
		return model.EmptyResponse{}, err
	}
	return model.EmptyResponse{}, nil
}
