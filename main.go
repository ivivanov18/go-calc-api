package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/ivivanov18/go-cal-api/handlers"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	router := http.NewServeMux()

	router.HandleFunc("GET /test", handlers.HandleTest)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	logger.Info("Server listening on port 8080")

	srv.ListenAndServe()
}
