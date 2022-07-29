package functionality

import (
	"github.com/google/uuid"
	"net-http-gorilla-ddd-example/models/entity"
	"net-http-gorilla-ddd-example/models/repository"
)

var _ BooksFunctionality = &Default{}

type BooksFunctionality interface {
	GetAll() ([]*entity.Book, error)
	Get(id uuid.UUID) (*entity.Book, error)
	Create(book *entity.Book) error
	Update(id uuid.UUID, book *entity.Book) error
	Delete(id uuid.UUID) error
}

type Default struct {
	Repository repository.BooksRepository
}

func (service *Default) GetAll() (b []*entity.Book, err error) {
	return service.Repository.GetAll()
}

func (service *Default) Get(id uuid.UUID) (b *entity.Book, err error) {
	return service.Repository.Get(id)
}

func (service *Default) Create(book *entity.Book) (err error) {
	return service.Repository.Create(book)
}

func (service *Default) Update(id uuid.UUID, book *entity.Book) (err error) {
	return service.Repository.Update(id, book)
}

func (service *Default) Delete(id uuid.UUID) (err error) {
	return service.Repository.Delete(id)
}
