package handler

import (
	"context"
	"net/http"

	"gitlab.ozon.dev/alexeyivashka/homework/checkout/internal/domain/model"
	"gitlab.ozon.dev/alexeyivashka/homework/libs/srvwrapper"
)

func (h *Handler) Purchase(w http.ResponseWriter, r *http.Request) {
	wrapper := srvwrapper.NewWrapper[*model.PurchaseRequest, *model.PurchaseResponse](h.handlePurchase)
	wrapper.ServeHTTP(w, r)
}

func (h *Handler) handlePurchase(ctx context.Context, request *model.PurchaseRequest) (*model.PurchaseResponse, error) {
	purchase, err := h.checkoutService.Purchase(ctx, request)
	if err != nil {
		return &model.PurchaseResponse{}, err
	}
	return &model.PurchaseResponse{OrderID: purchase.OrderID}, nil
}
