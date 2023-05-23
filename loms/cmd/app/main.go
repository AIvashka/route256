package main

import (
	"context"
	lg "gitlab.ozon.dev/alexeyivashka/homework/libs/logger"
	"gitlab.ozon.dev/alexeyivashka/homework/loms/cmd/app/config"
	h "gitlab.ozon.dev/alexeyivashka/homework/loms/internal/handler"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	logger := lg.GetLogger()

	cfg, err := config.LoadConfig()

	if err != nil {
		logger.Fatal("Error reading config file, %s", zap.Error(err))
	}

	handler := h.NewHandlerWithDependencies()
	h.SetupRoutes(handler)

	server := &http.Server{
		Addr: cfg.ServerAddress,
	}

	go func() {
		logger.Info("starting loms server")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("listen:%+s\n", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	logger.Info("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown: %s", zap.Error(err))
	}

	logger.Info("Server exiting")
}
