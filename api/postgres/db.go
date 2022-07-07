package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

type DB struct {
	*pgx.Conn
}

func Connect(url string) (*DB, error) {
	conn, err := pgx.Connect(context.Background(), url)

	if err != nil {
		return nil, err
	}

	if err := conn.Ping(context.Background()); err != nil {
		return nil, err
	}

	log.Println("Successfully connected to database")
	defer conn.Close(context.Background())
	return &DB{conn}, nil
}
