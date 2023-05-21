package handler

import (
	"context"
	"net/http"

	"gitlab.ozon.dev/alexeyivashka/homework/checkout/internal/domain/model"
	"gitlab.ozon.dev/alexeyivashka/homework/libs/srvwrapper"
)

func (h *Handler) DeleteFromCart(w http.ResponseWriter, r *http.Request) {
	wrapper := srvwrapper.NewWrapper[*model.DeleteFromCartRequest, *model.DeleteFromCartResponse](h.handleDeleteFromCartRequest)
	wrapper.ServeHTTP(w, r)
}

func (h *Handler) handleDeleteFromCartRequest(ctx context.Context, request *model.DeleteFromCartRequest) (*model.DeleteFromCartResponse, error) {
	response, err := h.checkoutService.DeleteFromCart(ctx, request)
	if err != nil {
		return &model.DeleteFromCartResponse{}, err
	}
	return response, nil
}
