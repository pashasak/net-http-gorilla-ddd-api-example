package api

import (
	"context"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net-http-gorilla-ddd-example/functionality"
	"net-http-gorilla-ddd-example/models/entity"
	"net-http-gorilla-ddd-example/utils"
	"net-http-gorilla-ddd-example/utils/log"
	"net/http"
	"time"
)

type Book struct {
	bookApp functionality.BooksFunctionality
	ctx     context.Context
}

func NewBook(ctx context.Context, bApp functionality.BooksFunctionality) *Book {
	return &Book{
		ctx:     ctx,
		bookApp: bApp,
	}
}

// GetBooks func gets all exists books.
// @Description Get all exists books.
// @Summary get all exists books
// @Tags Books
// @Accept json
// @Produce json
// @Success 200 {array} entity.Book
// @Router /v1/books [get]
func (bs *Book) GetBooks(w http.ResponseWriter, request *http.Request) {
	b, err := bs.bookApp.GetAll()
	if err != nil {
		utils.ErrorPresenter(w, err, http.StatusInternalServerError)
		return
		//log.Error(err.Error())
		//http.Error(w, "Ошибка запроса к БД.\n"+err.Error(), http.StatusInternalServerError)
	}
	utils.JsonPresenter(w, b)
	return
}

// GetBook func gets book by given ID or 404 error.
// @Description Get book by given ID.
// @Summary get book by given ID
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} entity.Book
// @Router /v1/book/{id} [get]
func (bs *Book) GetBook(w http.ResponseWriter, request *http.Request) {
	id, err := uuid.Parse(mux.Vars(request)["id"])
	if err != nil {
		log.Error(err.Error())
		http.Error(w, "Некорректный идентификатор. "+err.Error(), http.StatusBadRequest)
	}
	b, err := bs.bookApp.Get(id)
	utils.JsonPresenter(w, b)
	return
}

// UpdateBook func for updates book by given ID.
// @Description Update book.
// @Summary update book
// @Tags Book
// @Accept json
// @Produce json
// @Param id body string true "Book ID"
// @Param title body string true "Title"
// @Param author body string true "Author"
// @Param book_status body integer true "Book status"
// @Param book_attrs body entity.BookAttrs true "Book attributes"
// @Success 201 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/book [put]
func (bs *Book) UpdateBook(w http.ResponseWriter, request *http.Request) {
	id, err := uuid.Parse(mux.Vars(request)["id"])
	if err != nil {
		log.Log("error", err)
		http.Error(w, "Некорректный идентификатор. "+err.Error(), http.StatusBadRequest)
	}

	book := &entity.Book{}
	err = utils.JsonDecoder(request.Body, book)
	if err != nil {
		log.Log("error", err)
		http.Error(w, "Некорректная форма запроса", http.StatusBadRequest)
		return
	}

	// Create a new validator for a Book model.
	validate := utils.NewValidator()
	// Validate book fields.
	if err = validate.Struct(book); err != nil {
		// Return, if some fields are not valid.
		log.Log("error", err)
		http.Error(w, "Некорректные параметры запроса: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = bs.bookApp.Update(id, book)
	if err != nil {
		log.Log("error", "Ошибка записи", err)
		http.Error(w, "Некорректная форма запроса", http.StatusInternalServerError)
	}
	utils.JsonPresenter(w, id)
	return
}

// CreateBook func for creates a new book.
// @Description Create a new book.
// @Summary create a new book
// @Tags Book
// @Accept json
// @Produce json
// @Param title body string true "Title"
// @Param author body string true "Author"
// @Param book_attrs body entity.BookAttrs true "Book attributes"
// @Success 200 {object} entity.Book
// @Security ApiKeyAuth
// @Router /v1/book [post]
func (bs *Book) CreateBook(w http.ResponseWriter, request *http.Request) {

	b := &entity.Book{}

	err := utils.JsonDecoder(request.Body, &b)
	if err != nil {
		http.Error(w, "Некорректные параметры запроса: "+err.Error(), http.StatusBadRequest)
		return
	}
	// Set initialized default data for book:
	b.ID = uuid.New()
	b.CreatedAt = time.Now()
	b.BookStatus = 1 // 0 == draft, 1 == active

	// Create a new validator for a Book model.
	validate := utils.NewValidator()
	// Validate book fields.
	if err := validate.Struct(b); err != nil {
		// Return, if some fields are not valid.
		http.Error(w, "Некорректные параметры запроса: "+err.Error(), http.StatusBadRequest)
		return
	}
	err = bs.bookApp.Create(b)
	if err != nil {
		utils.ErrorPresenter(w, err, http.StatusInternalServerError)
		return
	}
	utils.JsonPresenter(w, b)
	return
}

// DeleteBook func for deletes book by given ID.
// @Description Delete book by given ID.
// @Summary delete book by given ID
// @Tags Book
// @Accept json
// @Produce json
// @Param id body string true "Book ID"
// @Success 204 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/book [delete]
func (bs *Book) DeleteBook(w http.ResponseWriter, request *http.Request) {
	id, err := uuid.Parse(mux.Vars(request)["id"])
	if err != nil {
		log.Log(err)
		http.Error(w, "Некорректный идентификатор. "+err.Error(), http.StatusBadRequest)
	}
	err = bs.bookApp.Delete(id)
	if err != nil {
		utils.ErrorPresenter(w, err, http.StatusInternalServerError)
		return
	}
	utils.JsonPresenter(w, id)
	return
}
