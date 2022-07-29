package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"net-http-gorilla-ddd-example/models/entity"
	"net-http-gorilla-ddd-example/models/repository"
)

type PGBooksRepository struct {
	conn *sqlx.DB
}

func NewBookRepository(db *sqlx.DB) repository.BooksRepository {
	return &PGBooksRepository{conn: db}
}

var _ repository.BooksRepository = &PGBooksRepository{}

func (q *PGBooksRepository) GetAll() (b []*entity.Book, err error) {

	// Define query string.
	query := `SELECT * FROM books`

	// Send query to database.
	err = q.conn.Select(&b, query)
	if err != nil {
		// Return empty object and error.
		return b, err
	}

	// Return query result.
	return b, err
}

func (q *PGBooksRepository) Get(id uuid.UUID) (b *entity.Book, err error) {
	// Define query string.
	query := `SELECT * FROM books WHERE id = $1`

	// Send query to database.
	err = q.conn.Get(&b, query, id)
	if err != nil {
		// Return empty object and error.
		return b, err
	}

	// Return query result.
	return b, err
}

func (q *PGBooksRepository) Create(b *entity.Book) error {
	// Define query string.
	query := `INSERT INTO books VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	// Send query to database.
	_, err := q.conn.Exec(query, b.ID, b.CreatedAt, b.UpdatedAt, b.UserID, b.Title, b.Author, b.BookStatus, b.BookAttrs)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

func (q *PGBooksRepository) Update(id uuid.UUID, b *entity.Book) error {
	query := `UPDATE books SET updated_at = $2, title = $3, author = $4, book_status = $5, book_attrs = $6 WHERE id = $1`

	// Send query to database.
	_, err := q.conn.Exec(query, id, b.UpdatedAt, b.Title, b.Author, b.BookStatus, b.BookAttrs)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

func (q *PGBooksRepository) Delete(id uuid.UUID) error {
	// Define query string.
	query := `DELETE FROM books WHERE id = $1`

	// Send query to database.
	_, err := q.conn.Exec(query, id)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}
