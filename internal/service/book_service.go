package service

import "github.com/salar-nourani/books-api/internal/repository"

type BookService struct {
	repo *repository.BookRepository
}

func NewBookService(r *repository.BookRepository) *BookService {
	return &BookService{repo: r}
}

func (s *BookService) GetAppStatus() string {
	return s.repo.GetBookStatus()
}

