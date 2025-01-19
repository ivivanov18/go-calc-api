package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/ivivanov18/go-cal-api/handlers"
	"github.com/ivivanov18/go-cal-api/handlers/add"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	router := http.NewServeMux()

	router.HandleFunc("GET /test", handlers.HandleTest)
	router.HandleFunc("POST /add", add.HandleAdd)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	logger.Info("Server listening on port 8080")

	srv.ListenAndServe()
}
