package main

import (
	"context"
	"log/slog"
	"main/views"
	"net/http"
	"os"
	"time"
)

func main() {
	router := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("public"))

	router.Handle("/public/*", http.StripPrefix("/public/", fileServer))
	// Move this to the router and handler package
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := views.Index()
		err := c.Render(r.Context(), w)
		if err != nil {
			http.Error(w, "Error rendering home template", http.StatusInternalServerError)
		}
	})

	killSig := make(chan os.Signal, 1)

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			slog.Error("server error", slog.Any("err", err))
		}
	}()

	slog.Info("Server started...")
	<-killSig

	slog.Info("Shutting down server")

	// Create a context with a timeout for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Server shutdown failed", slog.Any("err", err))
		os.Exit(1)
	}

	slog.Info("Server shutdown complete")
}
