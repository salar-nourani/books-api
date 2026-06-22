package main

import (
	"net/http"

	"github.com/salar-nourani/books-api/internal/handler"
	"github.com/salar-nourani/books-api/internal/repository"
	"github.com/salar-nourani/books-api/internal/service"
	"github.com/salar-nourani/books-api/pkg/config"
)

func main() {
	cfg := config.Load()
	logger := config.NewLogger()

	repo := repository.NewBookRepository()
	svc := service.NewBookService(repo)
	hdl := handler.NewBookHandler(svc)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		logger.Info().Msg("Health check endpoint hit")
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/status", hdl.StatusHandler)

	logger.Info().Msgf("Server starting on port %s", cfg.Port)
	http.ListenAndServe(":"+cfg.Port, nil)
}

