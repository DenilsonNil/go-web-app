package models

import (
	"log"
	"webapp/db"
)

type Product struct {
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
		log.Printf("Prouct: %s, Description: %s, Prive: %.2f, Quantity: %d\n", nome, descricao, preco, quantidade)
		produto := Product{nome, descricao, preco, quantidade}
		products = append(products, produto)
	}

	return products
}
