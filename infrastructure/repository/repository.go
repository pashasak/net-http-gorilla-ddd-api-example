package repository

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib" // load pgx driver for PostgreSQL
	"github.com/jmoiron/sqlx"
	"net-http-gorilla-ddd-example/config"
	"net-http-gorilla-ddd-example/models/repository"
	"net-http-gorilla-ddd-example/utils/log"
	"time"
)

type Repositories struct {
	Books repository.BooksRepository
	conn  *sqlx.DB
}

func NewRepositories() (*Repositories, error) {
	db, err := sqlx.Open("pgx", config.GetDBConnectString())
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxIdleConns(15)
	db.SetMaxOpenConns(30)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Repositories{
		Books: NewBookRepository(db),
		conn:  db,
	}, nil
}

//closes the  database connection
func (s *Repositories) Close() error {
	return s.conn.Close()
}

//This migrate all tables
func (s *Repositories) Automigrate() error {
	db, err := sql.Open("pgx", config.GetDBConnectString())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	driver, err := pgx.WithInstance(db, &pgx.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://platform/db/migrations",
		"postgres", driver)
	if err != nil {
		log.Fatal(err)
	}
	return m.Up()
}
