package main

import (
	"context"
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Produk struct {
	ID        int
	Name      string
	Price     int
	Stok      int
	Ket       string
	CreatedAt time.Time
	DeletedAt sql.NullTime
}

var DataProduk = []Produk{
	Produk{Name: "Iphone Duo", Price: 39000000},
	Produk{Name: "Samsung Fold 8", Price: 30000000},
}

// rupiah memformat angka jadi "Rp39.000.000"
func rupiah(n int) string {
	s := strconv.Itoa(n)
	hasil := ""
	for i, c := range s {
		// sisipkan titik tiap 3 digit dari kanan, kecuali setelah tanda minus
		if i > 0 && (len(s)-i)%3 == 0 && s[i-1] != '-' {
			hasil += "."
		}
		hasil += string(c)
	}
	return "Rp" + hasil
}

func add_one(n int) int {
	return n + 1
}

var funcs = template.FuncMap{"rupiah": rupiah, "add_one": add_one}

func parseHalaman(file string) *template.Template {
	return template.Must(template.New("layout").Funcs(funcs).ParseFiles(
		"templates/layout.html",
		"templates/header.html",
		"templates/sidebar.html",
		"templates/footer.html",
		"templates/"+file,
	))
}

var tplHome = parseHalaman("home.html")
var tplProfil = parseHalaman("profil.html")

var tplProduk = parseHalaman("produk/list.html")
var tplProdukForm = parseHalaman("produk/form.html")
var tplProdukDetail = parseHalaman("produk/detail.html")

func main() {
	_ = godotenv.Load() // baca .env jika ada

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN belum di-set (cek file .env)")
	}

	db, err := sql.Open("mysql", dsn) // db bertipe *sql.DB
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("database tidak bisa dihubungi: %v", err)
	}
	log.Println("database terhubung")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", homeHandler)
	mux.HandleFunc("GET /profil", profilHandler)
	mux.HandleFunc("GET /produk", produkHandler(db))
	mux.HandleFunc("GET /produk/form", produkFormHandler)
	mux.HandleFunc("POST /produk", produkSimpanHandler(db))
	mux.HandleFunc("GET /produk/{id}", produkByIdHandler(db))
	mux.HandleFunc("POST /produk/{id}/hapus", produkHapusHandler(db))

	log.Fatal(http.ListenAndServe(":8081", mux))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplHome.ExecuteTemplate(w, "layout", nil)
}

func profilHandler(w http.ResponseWriter, r *http.Request) {
	tplProfil.ExecuteTemplate(w, "layout", nil)
}

func produkHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		produk, err := ambilSemuaProduk(r.Context(), db)

		if err != nil {
			log.Fatal(err)
		}
		tplProduk.ExecuteTemplate(w, "layout", produk)
	}
}

func produkFormHandler(w http.ResponseWriter, r *http.Request) {
	tplProdukForm.ExecuteTemplate(w, "layout", nil)
}

func produkByIdHandler(db *sql.DB) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		param_id, _ := strconv.Atoi(id)
		//ambil data dari database
		produk, err := getProdukById(r.Context(),db, param_id)
		if err != nil{
			log.Fatal(err)
		}
		tplProdukDetail.ExecuteTemplate(w, "layout", produk)
	}
}

func getProdukById(ctx context.Context, db *sql.DB, id int) (Produk, error) {
    var p Produk
    err := db.QueryRowContext(ctx,
        "SELECT id, nama, harga, stok, COALESCE(keterangan,'') AS keterangan, created_at, deleted_at FROM produk WHERE id = ? AND deleted_at IS NULL",
        id,
    ).Scan(&p.ID, &p.Name, &p.Price, &p.Stok, &p.Ket, &p.CreatedAt, &p.DeletedAt)
    if err != nil {
        return Produk{}, err
    }
    return p, nil
}

func produkSimpanHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		price, err := strconv.Atoi(r.FormValue("price"))
		if name == "" || err != nil {
			http.Error(w, "Nama dan harga wajib diisi dengan benar", http.StatusBadRequest)
			return
		}
		stok, _ := strconv.Atoi(r.FormValue("stok")) // kosong -> 0

		p := Produk{
			Name:      name,
			Price:     price,
			Stok:      stok,
			Ket:       r.FormValue("ket"),
			CreatedAt: time.Now(),
		}
		if _, err := simpanProduk(r.Context(), db, p); err != nil {
			log.Println("gagal simpan produk:", err)
			http.Error(w, "Gagal menyimpan produk", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/produk", http.StatusSeeOther)
	}
}
func simpanProduk(ctx context.Context, db *sql.DB, p Produk) (int64, error){
	res, err := db.ExecContext(ctx,
		"INSERT INTO produk (nama, harga, stok, keterangan, created_at) VALUES (?, ? ,? ,?, ?)",
		p.Name, p.Price, p.Stok, p.Ket, p.CreatedAt,
	)
	if err != nil{
		return 0, err
	}
	id, err := res.LastInsertId()
	return id, err
}

func ambilSemuaProduk(ctx context.Context, db *sql.DB) ([]Produk, error) {
    rows, err := db.QueryContext(ctx,
        "SELECT id, nama, harga, stok, coalesce(keterangan,'') as keterangan, created_at, deleted_at FROM produk WHERE deleted_at IS NULL ORDER BY id")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var hasil []Produk
    for rows.Next() {
        var p Produk
        if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stok, &p.Ket, &p.CreatedAt, &p.DeletedAt); err != nil {
            return nil, err
        }
        hasil = append(hasil, p)
    }
    return hasil, rows.Err()
}

func hapusProduk(ctx context.Context, db *sql.DB, id int) error {
    _, err := db.ExecContext(ctx, "UPDATE produk SET deleted_at = NOW() WHERE id = ? AND deleted_at IS NULL", id)
    return err
}

func produkHapusHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        idStr := r.PathValue("id")
        id, err := strconv.Atoi(idStr)
        if err != nil {
            http.Error(w, "ID tidak valid", http.StatusBadRequest)
            return
        }

        if err := hapusProduk(r.Context(), db, id); err != nil {
            log.Println("Gagal hapus produk:", err)
            http.Error(w, "Gagal menghapus produk", http.StatusInternalServerError)
            return
        }
        
        // Redirect kembali ke daftar produk setelah sukses hapus
        http.Redirect(w, r, "/produk", http.StatusSeeOther)
    }
}