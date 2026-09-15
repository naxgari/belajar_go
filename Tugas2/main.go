/*
- Server dengan minimal 5 route, di antaranya:
- satu route memakai path parameter `{id}`
- satu route memakai query string dengan nilai default
- satu route `POST` yang membalas `201 Created`
- halaman utama memakai `{$}`
- Submit: file `.go` + screenshot output `curl -i` untuk tiap route, dalam satu zip
*/

package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", homeHandler)
	mux.HandleFunc("GET /about/", aboutHandler)
	mux.HandleFunc("GET /halo/{nama...}", haloHandler)
	mux.HandleFunc("GET /produk/{id}", detailProduk)
	mux.HandleFunc("GET /produk", produkHandler)
	mux.HandleFunc("POST /produk", tambahProduk)
	
	fmt.Println("Server berjalan di http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Selamat datang di halaman utama!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Ini adalah halaman tentang kami.")
}

func haloHandler(w http.ResponseWriter, r *http.Request) {
	nama := r.PathValue("nama")
	fmt.Fprintf(w, "Halo, %s!\n", nama)
}

func detailProduk(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "Detail Produk ID: %s\n", id)
}

func produkHandler(w http.ResponseWriter, r *http.Request) {
	kategori := r.URL.Query().Get("kategori")
	if kategori == "" {
		kategori = "semua"
	}
	fmt.Fprintf(w, "Daftar Produk - Kategori: %s\n", kategori)
}

func tambahProduk(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Produk Baru Berhasil Ditambahkan!")
}