package repository

import (
	"github.com/google/uuid"
	"net-http-gorilla-ddd-example/models/entity"
)

type BooksRepository interface {
	GetAll() ([]*entity.Book, error)
	Get(id uuid.UUID) (*entity.Book, error)
	Create(book *entity.Book) error
	Update(id uuid.UUID, book *entity.Book) error
	Delete(id uuid.UUID) error
}
