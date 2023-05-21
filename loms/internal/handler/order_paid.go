package handler

import (
	"context"
	"net/http"

	"gitlab.ozon.dev/alexeyivashka/homework/libs/srvwrapper"
	"gitlab.ozon.dev/alexeyivashka/homework/loms/internal/domain/model"
)

func (h *Handler) OrderPaid(w http.ResponseWriter, r *http.Request) {
	wrapper := srvwrapper.NewWrapper[model.OrderPaidRequest, model.EmptyResponse](h.handleOrderPaidRequest)
	wrapper.ServeHTTP(w, r)
}

func (h *Handler) handleOrderPaidRequest(ctx context.Context, request model.OrderPaidRequest) (model.EmptyResponse, error) {
	err := h.orderPaidUseCase.Execute(ctx, request.ID)
	if err != nil {
		return model.EmptyResponse{}, err
	}
	return model.EmptyResponse{}, nil
}
