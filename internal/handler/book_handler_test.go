package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/salar-nourani/books-api/internal/repository"
	"github.com/salar-nourani/books-api/internal/service"
)

func TestStatusHandler(t *testing.T) {
	repo := repository.NewBookRepository()
	svc := service.NewBookService(repo)
	hdl := NewBookHandler(svc)

	req := httptest.NewRequest(http.MethodGet, "/status", nil)
	rec := httptest.NewRecorder()

	hdl.StatusHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, rec.Code)
	}

	want := "Database is connected and healthy"
	got := rec.Body.String()

	if got != want {
		t.Errorf("expected body %q, got %q", want, got)
	}
}

