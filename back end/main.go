package main

import (
	"backend/api"
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {

	db, err := sql.Open(
		"pgx",
		"postgres://nurseapp:%20@localhost:5432/nurse_call",
	)

	if err != nil {
		log.Fatal(err)
	}

	server := api.NewServer(":3000", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}
