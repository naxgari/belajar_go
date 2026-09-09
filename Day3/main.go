package main

//package net/http
//server

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	//GET /{$} jika kita isikan acak maka akan memunculkan 404 not found, karena tidak ada route yang sesuai, jika kita isikan /about maka akan memunculkan about page
	mux.HandleFunc("GET /{$}", homeHandler) //nanti kita buat fungsi homeHandlernya, handler itu khusus untuk mengangani satu route, kalau ada banyak route bisa dibuat banyak handler
	mux.HandleFunc("GET /about/", aboutHandler)
	mux.HandleFunc("GET /halo/{nama...}", haloHandler) //berbeda dengan {nama}, {nama...} itu bisa menampung banyak parameter, misal /halo/king/123/456, maka yang ditangkap hanya king, 123, 456 tidak akan ditangkap 
	
	mux.HandleFunc("GET /produk/{id}", detailProduk)
	mux.HandleFunc("GET /produk/baru", formProduk)

	/*
	ada yang tidak boleh dalam routing
	mux.HandleFunc("GET /a/{x}", handlea)
	mux.HandleFunc("GET /{y}/b", handleb)
	*/

	http.ListenAndServe(":8080", mux)
}

func homeHandler(w http.ResponseWriter, r *http.Request) { //ada 2 parameter, w untuk menulis response, r untuk membaca request
	w.Write([]byte("Hello World\n")) //untuk menampilkan ke web browser menggunakan Write
	fmt.Println("Request received") //untuk menampilkan ke terminal
	fmt.Fprint(w, "Kelas King") //untuk menampilkan ke web browser menggunakan Fprint
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //untuk mengatur header response, misal content type json
	w.WriteHeader(http.StatusCreated) //untuk mengatur status code response, misal 200 OK
	path_url := r.URL.Path //untuk mengambil path url yang diminta
	mode := r.URL.Query().Get("mode") //untuk mengambil method request yang diminta
	page := r.URL.Query().Get("page") //untuk mengambil query parameter yang diminta
	w.Write([]byte("About Page\n"))
	fmt.Fprint(w, "Path URL:", path_url, "\n")
	fmt.Fprint(w, "Mode:", mode, "\n")
	fmt.Fprint(w, "Page:", page, "\n")
	fmt.Println("Request received to about page")
	fmt.Fprint(w, "Ini page about king")
}

func haloHandler(w http.ResponseWriter, r *http.Request) {
	nama := r.PathValue("nama")
	fmt.Println("Request received to halo page")
	fmt.Fprint(w, "Kelas King " +nama)

}

/*
Request received to about page
Request received

saat memasuki halaman about, maka akan menampilkan "Request received to about page" di terminal, 
dan saat memasuki halaman home, maka akan menampilkan "Request received" di terminal
*/

func detailProduk(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Println("Request received to detail produk page")
	fmt.Fprint(w, "Detail Produk dengan ID: " +id)
}

func formProduk(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received to form produk page")
	fmt.Fprint(w, "Form Produk Baru")
}