package db

import (
	"database/sql"
	"log"
)

func DbConnect() *sql.DB {
	connStr := "user=root dbname=loja password=root sslmode=disable host=localhost port=5432"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Conexão com o banco de dados estabelecida com sucesso!")
	}

	return db
}
