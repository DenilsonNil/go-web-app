package models

import (
	"fmt"
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
	rows, err := db.Query("SELECT * FROM produtos order by id asc")
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

func GetProductByID(id string) Product {
	db := db.DbConnect()
	row := db.QueryRow("SELECT * FROM produtos WHERE id=$1", id)

	var produto Product
	err := row.Scan(&produto.Id, &produto.Nome, &produto.Descricao, &produto.Preco, &produto.Quantidade)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	fmt.Println("Produto buscado: ", produto)
	return produto
}

func UpdateProduct(product Product) {
	db := db.DbConnect()
	updateStatement := "UPDATE produtos SET nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5"
	_, err := db.Exec(updateStatement, product.Nome, product.Descricao, product.Preco, product.Quantidade, product.Id)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Produto atualizado com sucesso:", product.Nome)
	defer db.Close()
}
