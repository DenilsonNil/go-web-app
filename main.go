package main

import (
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Notebook", Descricao: "Notebook Dell", Preco: 4500.00, Quantidade: 10},
		{Nome: "Smartphone", Descricao: "Smartphone Samsung", Preco: 2500.00, Quantidade: 15},
		{Nome: "Tablet", Descricao: "Tablet Apple", Preco: 3500.00, Quantidade: 5},
		{Nome: "Sapato de couro de jacaré", Descricao: "Organizações Tabajara", Preco: 1.00, Quantidade: 7},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}
