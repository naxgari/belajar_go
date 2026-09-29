package main

import (
	"html/template"
	"net/http"
	"strconv"
)

type PageProduk struct{
	Title string
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

func parseHalaman(file string) *template.Template{
	return template.Must(template.New("layout").Funcs(funcs).ParseFiles(
		"templates/layout.html",
		"templates/header.html",
		"templates/sidebar.html",
		"templates/footer.html",
		"templates/"+file,
	))
}

var tplProfil = parseHalaman("profil.html")
var tplProduk = parseHalaman("produk.html")

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", homeHandler)
	mux.HandleFunc("GET /profil", profilHandler)
	mux.HandleFunc("GET /produk", produkHandler)

	http.ListenAndServe(":8080", mux)
}

func profilHandler(w http.ResponseWriter, r *http.Request){
	tplProfil.ExecuteTemplate(w, "layout", nil)
}

func produkHandler(w http.ResponseWriter, r *http.Request){
	//data := PageProduk{Title: "Produk", Produk: DataProduk}
	tplProduk.ExecuteTemplate(w, "layout", DataProduk)
}

func homeHandler(w http.ResponseWriter, r *http.Request){
	
}
