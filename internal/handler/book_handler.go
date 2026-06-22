package handler

import (
	"net/http"
	"github.com/salar-nourani/books-api/internal/service"
)

type BookHandler struct {
	svc *service.BookService
}

func NewBookHandler(s *service.BookService) *BookHandler {
	return &BookHandler{svc: s}
}

func (h *BookHandler) StatusHandler(w http.ResponseWriter, r *http.Request) {
	status := h.svc.GetAppStatus()
	w.Write([]byte(status))
}

