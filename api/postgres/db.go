package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DB struct {
	*pgxpool.Pool
}

func Connect(url string) (*DB, error) {
	dbpool, err := pgxpool.Connect(context.Background(), url)

	if err != nil {
		return nil, err
	}

	if err := dbpool.Ping(context.Background()); err != nil {
		return nil, err
	}

	log.Println("Successfully connected to database")
	defer dbpool.Close()
	return &DB{dbpool}, nil
}
