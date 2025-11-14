package controllers

import (
	"fmt"
	"net/http"
	"text/template"
	"webapp/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	allProducts := models.GetAllProducts()
	fmt.Println("All products:", allProducts)

	// produtos := []Produto{
	// 	{Nome: "Notebook", Descricao: "Notebook Dell", Preco: 4500.00, Quantidade: 10},
	// 	{Nome: "Smartphone", Descricao: "Smartphone Samsung", Preco: 2500.00, Quantidade: 15},
	// 	{Nome: "Tablet", Descricao: "Tablet Apple", Preco: 3500.00, Quantidade: 5},
	// 	{Nome: "Sapato de couro de jacaré", Descricao: "Organizações Tabajara", Preco: 1.00, Quantidade: 7},
	// }

	temp.ExecuteTemplate(w, "Index", allProducts)
}
