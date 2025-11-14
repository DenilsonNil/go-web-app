package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

type Product struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func dbConnect() *sql.DB {
	connStr := "user=root dbname=loja password=root sslmode=disable host=localhost port=5432"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Conexão com o banco de dados estabelecida com sucesso!")
	}

	return db
}

func getAllProducts(db *sql.DB) []Product {
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
func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := dbConnect()
	allProducts := getAllProducts(db)
	fmt.Println("All products:", allProducts)

	// produtos := []Produto{
	// 	{Nome: "Notebook", Descricao: "Notebook Dell", Preco: 4500.00, Quantidade: 10},
	// 	{Nome: "Smartphone", Descricao: "Smartphone Samsung", Preco: 2500.00, Quantidade: 15},
	// 	{Nome: "Tablet", Descricao: "Tablet Apple", Preco: 3500.00, Quantidade: 5},
	// 	{Nome: "Sapato de couro de jacaré", Descricao: "Organizações Tabajara", Preco: 1.00, Quantidade: 7},
	// }

	temp.ExecuteTemplate(w, "Index", allProducts)
}
