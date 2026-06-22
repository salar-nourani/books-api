package repository

type BookRepository struct{}

func NewBookRepository() *BookRepository {
	return &BookRepository{}
}

func (r *BookRepository) GetBookStatus() string {
	return "Database is connected and healthy"
}

