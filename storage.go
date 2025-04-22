package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	Init() error // init the db with required tables
	GetLongURL(shortURL string) (string, error)
	AddShortURL(longURL string) (string, error)
}

type PGStorage struct {
	db *sql.DB
}

func NewStorage() (*PGStorage, error) {
	connStr := "postgres://postgres:verystrongpassword@localhost:8080/us_db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PGStorage{
		db: db,
	}, nil
}

func (pg *PGStorage) Init() error {
	query := `create table if not exists url (
		id serial primary key,
		shortURL varchar(100),
		LongURL varchar(100),
	)`

	_, err := pg.db.Exec(query)
	return err
}

func (pg *PGStorage) GetLongURL(shortURL string) (string, error) {
	return "", nil
}

func (pg *PGStorage) AddShortURL(longURL string) (string, error) {
	return "", nil
}
