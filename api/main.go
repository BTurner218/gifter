package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/BTurner218/gifter/api/server"
	"github.com/go-chi/jwtauth/v5"
	"github.com/jackc/pgx/v4"
)

func main() {
	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": 123})
	fmt.Printf(tokenString)

	conn, err := pgx.Connect(context.Background(), os.Getenv("POSTGRESQL_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	srv := server.NewServer()
	log.Fatal(srv.Run())
}
