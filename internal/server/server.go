package server

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	gracefulShutdownTimeout = 30 * time.Second
)

func Run(handler http.Handler, addr string) error {
	logger := slog.Default()

	srv := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	errCh := make(chan error, 1)
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 'gracefulShutdownTimeout' seconds.
	quit := make(chan os.Signal, 1)

	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")

	select {
	case err := <-errCh:
		return err
	case <-quit:
		// The context is used to inform the server it has 'gracefulShutdownTimeout' seconds to finish
		// the request it is currently handling
		ctx, cancel := context.WithTimeout(context.Background(), gracefulShutdownTimeout)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			logger.Error("Server forced to shutdown", slog.String("error", err.Error()))
			return err
		}
		logger.Info("Server exited gracefully")
		return nil
	default:
		return srv.Close()
	}
}
