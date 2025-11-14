package models

import (
	"log"
	"webapp/db"
)

type Product struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func GetAllProducts() []Product {
	db := db.DbConnect()
	rows, err := db.Query("SELECT * FROM produtos")
	products := []Product{}
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var nome string
		var descricao string
		var preco float64
		var quantidade int

		if err := rows.Scan(&id, &nome, &descricao, &preco, &quantidade); err != nil {
			log.Fatal(err)
		}

		produto := Product{id, nome, descricao, preco, quantidade}
		products = append(products, produto)
	}

	return products
}

func InsertProduct(product Product) {
	db := db.DbConnect()
	insertStatement := "INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES ($1, $2, $3, $4)"
	_, err := db.Exec(insertStatement, product.Nome, product.Descricao, product.Preco, product.Quantidade)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Produto inserido com sucesso:", product.Nome)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.DbConnect()
	deleteStatement := "DELETE FROM produtos WHERE id=$1"
	_, err := db.Exec(deleteStatement, id)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Produto deletado com sucesso, ID:", id)

	defer db.Close()
}
