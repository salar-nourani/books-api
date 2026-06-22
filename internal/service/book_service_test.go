package service

import (
	"testing"

	"github.com/salar-nourani/books-api/internal/repository"
)

func TestGetAppStatus(t *testing.T) {
	repo := repository.NewBookRepository()
	svc := NewBookService(repo)

	got := svc.GetAppStatus()
	want := "Database is connected and healthy"

	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

