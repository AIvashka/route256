package handler

import (
	"context"
	"net/http"

	"gitlab.ozon.dev/alexeyivashka/homework/checkout/internal/domain/model"
	"gitlab.ozon.dev/alexeyivashka/homework/libs/srvwrapper"
)

func (h *Handler) AddToCart(w http.ResponseWriter, r *http.Request) {
	wrapper := srvwrapper.NewWrapper[*model.AddToCartRequest, *model.AddToCartResponse](h.handleAddToCart)
	wrapper.ServeHTTP(w, r)
}

func (h *Handler) handleAddToCart(ctx context.Context, request *model.AddToCartRequest) (*model.AddToCartResponse, error) {
	response, err := h.checkoutService.AddToCart(ctx, request)
	if err != nil {
		return &model.AddToCartResponse{}, err
	}
	return response, nil
}
