package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/ivivanov18/go-cal-api/handlers"
	"github.com/ivivanov18/go-cal-api/handlers/add"
	"github.com/ivivanov18/go-cal-api/handlers/multiply"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	router := http.NewServeMux()

	router.HandleFunc("GET /test", handlers.HandleTest)
	router.HandleFunc("POST /add", add.Handle)
	router.HandleFunc("POST /multiply", multiply.Handle)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	logger.Info("Server listening on port 8080")

	srv.ListenAndServe()
}
