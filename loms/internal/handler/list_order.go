package handler

import (
	"context"
	"net/http"

	"gitlab.ozon.dev/alexeyivashka/homework/libs/srvwrapper"
	"gitlab.ozon.dev/alexeyivashka/homework/loms/internal/domain/model"
)

func (h *Handler) ListOrder(w http.ResponseWriter, r *http.Request) {
	wrapper := srvwrapper.NewWrapper[model.ListOrderRequest, model.ListOrderResponse](h.handleListOrderRequest)
	wrapper.ServeHTTP(w, r)
}

func (h *Handler) handleListOrderRequest(ctx context.Context, request model.ListOrderRequest) (model.ListOrderResponse, error) {
	order, err := h.listOrderUseCase.Execute(ctx, request.ID)
	if err != nil {
		return model.ListOrderResponse{}, err
	}
	return model.ListOrderResponse{Status: order.Status,
		UserID:     order.UserID,
		OrderItems: order.Items,
	}, nil
}
