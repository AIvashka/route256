package handler

import (
	"net/http"

	"gitlab.ozon.dev/alexeyivashka/homework/checkout/internal/domain/model"
)

type Handler struct {
	checkoutService model.CheckoutService
}

func NewHandler(cs model.CheckoutService) *Handler {
	return &Handler{checkoutService: cs}
}

func SetupRoutes(handler *Handler) {
	http.HandleFunc("/purchase", handler.Purchase)
	http.HandleFunc("/listCart", handler.ListCart)
	http.HandleFunc("/deleteFromCart", handler.DeleteFromCart)
	http.HandleFunc("/addToCart", handler.AddToCart)
}
