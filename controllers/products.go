package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"webapp/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.GetAllProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}

// Load the form to create a new product
func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	log.Println("Method:", r.Method)
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			fmt.Println("Erro na conversão do preço:", err)
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			fmt.Println("Erro na conversão da quantidade:", err)
		}

		produto := models.Product{
			Nome:       nome,
			Descricao:  descricao,
			Preco:      precoConvertido,
			Quantidade: quantidadeConvertida,
		}

		models.InsertProduct(produto)
		http.Redirect(w, r, "/", 301)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	models.DeleteProduct(idProduto)
	http.Redirect(w, r, "/", 301)
}
