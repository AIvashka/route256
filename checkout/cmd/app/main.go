package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gitlab.ozon.dev/alexeyivashka/homework/checkout/cmd/app/config"
	"gitlab.ozon.dev/alexeyivashka/homework/checkout/internal/app/checkout"
	h "gitlab.ozon.dev/alexeyivashka/homework/checkout/internal/handler"
	"gitlab.ozon.dev/alexeyivashka/homework/checkout/internal/storage"
	lg "gitlab.ozon.dev/alexeyivashka/homework/libs/logger"
	"gitlab.ozon.dev/alexeyivashka/homework/libs/loms_client"
	"gitlab.ozon.dev/alexeyivashka/homework/libs/product_service_client"
	"go.uber.org/zap"
)

func main() {

	logger := lg.GetLogger()

	cfg, err := config.LoadConfig(true)

	if err != nil {
		logger.Fatal("Error reading config file, %s", zap.Error(err))
	}

	fmt.Print(cfg)

	lomsClient := loms_client.NewClient(cfg.LomsUrl)
	mockCartRepository := storage.NewMockCartRepository()
	productService := product_service_client.NewClient(cfg.ProductServiceUrl, cfg.ProductServiceToken)

	cs := checkout.NewService(lomsClient, mockCartRepository, productService)

	handler := h.NewHandler(cs)
	h.SetupRoutes(handler)

	server := &http.Server{
		Addr: ":8081",
	}

	go func() {
		logger.Info("starting checkout server")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("listen:%+s\n", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	logger.Info("checkout server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("checkout server forced to shutdown: %s", zap.Error(err))
	}

	logger.Info("checkout server exiting")
}
