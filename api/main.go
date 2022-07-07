package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BTurner218/gifter/api/postgres"
	"github.com/BTurner218/gifter/api/server"
	"github.com/go-chi/jwtauth/v5"
)

func main() {
	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": 123})
	fmt.Println(tokenString)

	_, err := postgres.Connect(os.Getenv("POSTGRESQL_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	srv := server.NewServer()
	log.Fatal(srv.Run())
}
