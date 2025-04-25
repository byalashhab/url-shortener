package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Storage interface {
	Init() error // init the db with required tables
	AddShortURL(longURL string) (string, error)
	GetLongURL(shortURL string) (string, error)
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
	query := `create table if not exists url(
		id int primary key,
		shortURL varchar(100),
		longURL varchar(100)
	)`

	_, err := pg.db.Exec(query)
	return err
}

func (pg *PGStorage) AddShortURL(longURL string) (string, error) {

	var shortURL string
	q := "select shortURL from url where longURL = $1"

	err := pg.db.QueryRow(q, longURL).Scan(&shortURL)

	if err != nil {
		log.Fatal(err)
	}

	if shortURL != "" {
		return shortURL, nil
	}

	hash, id := GenerateHash(longURL)

	q = "insert into url values($1, $2, $3)"

	_, err = pg.db.Exec(q, id, hash, longURL)

	if err != nil {
		log.Fatal(err)
	}

	return hash, nil
}

func (pg *PGStorage) GetLongURL(shortURL string) (string, error) {

	var longURL string
	q := "select longURL from url where shortURL = $1"

	err := pg.db.QueryRow(q, shortURL).Scan(&longURL)

	if err != nil {
		log.Fatal(err)
	}

	return longURL, nil
}
