package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v4"
)

func main() {
	// TODO: Use environment variable for database URL
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:example@localhost:5432/gifter")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", hello)
	http.ListenAndServe(":8080", r)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
