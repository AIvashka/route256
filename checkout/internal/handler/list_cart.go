package handler

import (
	"context"
	"net/http"

	"gitlab.ozon.dev/alexeyivashka/homework/checkout/internal/domain/model"
	"gitlab.ozon.dev/alexeyivashka/homework/libs/srvwrapper"
)

func (h *Handler) ListCart(w http.ResponseWriter, r *http.Request) {
	wrapper := srvwrapper.NewWrapper[*model.ListCartRequest, *model.ListCartResponse](h.handleListCart)
	wrapper.ServeHTTP(w, r)
}

func (h *Handler) handleListCart(ctx context.Context, request *model.ListCartRequest) (*model.ListCartResponse, error) {
	response, err := h.checkoutService.ListCart(ctx, request)
	if err != nil {
		return &model.ListCartResponse{}, err
	}
	return response, nil
}
