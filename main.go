package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Args, os.Getenv, os.Stdin, os.Stdout, os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(
	ctx context.Context,
	args []string,
	getenv func(string) string,
	stdin io.Reader,
	stdout, stderr io.Writer,
) error {
	interuptCtx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()
	baseCtx, stop := context.WithCancel(ctx)
	defer stop()
	srv := &http.Server{
		Addr:    ":7443",
		Handler: newHandler(),
		BaseContext: func(net.Listener) context.Context {
			return baseCtx
		},
	}
	go func() {
		slog.Info("Server started!", "port", 7443)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	// wait for SIGINT/SIGTERM signal to start graceful shutdown procedure
	<-interuptCtx.Done()
	slog.Info("Interrupted signal received. Gracefully shutting down server....")
	cancel() // instantly stop the application on further interrupt signals
	stop()   // signal all handlers that the server is shutting down
	// new context needed to force shutdown after timeout
	shutdownCtx, shutdown := context.WithTimeout(ctx, 15*time.Second)
	defer shutdown()
	return srv.Shutdown(shutdownCtx)
}

func newHandler() http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux)
	return mux
}
