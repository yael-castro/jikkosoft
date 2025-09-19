package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/yael-castro/jikkosoft/3/internal/container"
	"github.com/yael-castro/jikkosoft/3/internal/runtime"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Setting exit code
	exitCode := 0
	defer func() {
		os.Exit(exitCode)
	}()

	// Building main context
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGTERM)
	defer stop()

	// Setting a default logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	slog.SetDefault(logger) // Setting default logger

	// Initializing *echo.Echo server
	var e *echo.Echo

	c := container.New()

	if err := c.Inject(ctx, &e); err != nil {
		slog.ErrorContext(ctx, "failed_server_built", "error", err)
		return
	}

	// Getting http port
	port := os.Getenv("PORT")
	if len(port) == 0 {
		const defaultPort = "8080"
		port = defaultPort
	}

	// Listening for shutdown gracefully
	shutdownCh := make(chan struct{}, 1)

	go func() {
		defer close(shutdownCh)

		<-ctx.Done()
		shutdown(e)

		shutdownCh <- struct{}{}
	}()

	// Running http server
	errCh := make(chan error, 1)

	go func() {
		defer close(errCh)

		slog.InfoContext(ctx, "running", "version", runtime.GitCommit, "port", port)
		errCh <- e.Start(":" + port)
	}()

	// Waiting for cancellation or error
	var err error

	select {
	case <-ctx.Done():
	case err = <-errCh:
		exitCode = 1
	}

	stop()
	<-shutdownCh

	slog.Error("exit", "code", exitCode, "error", err)
}

func shutdown(e *echo.Echo) {
	const gracePeriod = 10 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), gracePeriod)
	defer cancel()

	// Closing http server
	err := e.Shutdown(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "failed_shutdown", "error", err)
		return
	}

	slog.InfoContext(ctx, "success_shutdown")
}
