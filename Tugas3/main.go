package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Produk struct {
	ID    int
	Nama  string
	Harga int
}

var (
	produkStore = []Produk{
		{ID: 1, Nama: "Kipas Panasonic", Harga: 450000},
		{ID: 2, Nama: "Kipas Super Ultrasonic", Harga: 250000000},
	}
	nextID = 3
	mu     sync.Mutex
)

func renderTemplate(w http.ResponseWriter, page string, data map[string]interface{}) {
	files := []string{
		"templates/header.html",
		"templates/footer.html",
		"templates/sidebar.html",
		"templates/layout.html",
		"templates/profil.html",
		"templates/produk.html",
		"templates/produk_form.html",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Gagal memuat template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if data == nil {
		data = make(map[string]interface{})
	}
	data["CurrentPage"] = page

	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, "Gagal merender template: "+err.Error(), http.StatusInternalServerError)
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	renderTemplate(w, "profil", map[string]interface{}{
		"Title": "Beranda / Profil",
	})
}

func handleProfil(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "profil", map[string]interface{}{
		"Title": "Profil Pengguna",
	})
}

func handleProdukList(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	renderTemplate(w, "produk", map[string]interface{}{
		"Title":  "Daftar Produk",
		"Produk": produkStore,
	})
}

func handleProdukForm(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "produk_form", map[string]interface{}{
		"Title": "Tambah Produk Baru",
	})
}

func handleProdukSimpan(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/produk", http.StatusSeeOther)
		return
	}

	nama := r.FormValue("nama")
	hargaStr := r.FormValue("harga")
	harga, _ := strconv.Atoi(hargaStr)

	if nama != "" && harga > 0 {
		mu.Lock()
		produkStore = append(produkStore, Produk{
			ID:    nextID,
			Nama:  nama,
			Harga: harga,
		})
		nextID++
		mu.Unlock()
	}

	http.Redirect(w, r, "/produk", http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/profil", handleProfil)
	http.HandleFunc("/produk", handleProdukList)
	http.HandleFunc("/produk/form", handleProdukForm)
	http.HandleFunc("/produk/simpan", handleProdukSimpan)

	log.Println("Server berjalan di port :8080 (http://localhost:8080)...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}