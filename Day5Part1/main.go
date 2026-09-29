package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Homepage struct{
	Title string
	MaxLoop int
	Data []string
	Produk []Produk
}

type Produk struct{
	Name	string
	Price	int
}

var DataProduk = []Produk{
	Produk{Name: "Iphone Duo", Price: 39000000},
	Produk{Name: "Samsung Fold 8", Price: 30000000},
}

func rupiah(n int) string{
	return "Rp" + strconv.Itoa(n)
}
var funcs = template.FuncMap{"rupiah": rupiah}
var tpl = template.Must(template.New("layout").Funcs(funcs).ParseFiles("templates/layout.html"))

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", homeHandler)

	http.ListenAndServe(":8080", mux)
}

func homeHandler(w http.ResponseWriter, r *http.Request){
	data_homepage := Homepage{Title: "Homepage", MaxLoop: 1, Data:[]string{"Kerbau", "Sapi", "Kuda"}, Produk: DataProduk}
	
	if err := tpl.Execute(w, data_homepage); err != nil{
		log.Printf("render halaman gagal: %v", err)
	}
}
