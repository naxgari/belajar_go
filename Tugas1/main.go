/*Buat program "Daftar Belanja" dengan:
- struct Item (Nama, Harga, Jumlah)
- Slice untuk menyimpan item
- Fungsi untuk tambah, hitung total, tampilkan
- Validasi: harga dan jumlah harus > 0

File yang dikumpulkan:
- main.go+ screenshot output dalam zip*/

package main

import "fmt"

type Item struct { // struct item belanja
	Nama   string
	Harga  float64
	Jumlah int
}

var daftarBelanja []Item // slice untuk nyimpen item belanja

func tambahItem(nama string, harga float64, jumlah int) {
	if harga <= 0 && jumlah <= 0 { // validasi harga dan jumlah
		fmt.Printf("*Warning: Harga dan jumlah %s harus lebih dari 0\n\n", nama)
		return
	}
	if harga <= 0 { // validasi harga
		fmt.Printf("*Warning: Harga %s harus lebih dari 0\n\n", nama)
		return
	}
	if jumlah <= 0 { // validasi jumlah
		fmt.Printf("*Warning: Jumlah %s harus lebih dari 0\n\n", nama)
		return
	}

	daftarBelanja = append(daftarBelanja, Item{
		Nama:   nama,
		Harga:  harga,
		Jumlah: jumlah,
	})
}

func hitungTotal() float64 {
	var total float64
	for _, item := range daftarBelanja {
		total += item.Harga * float64(item.Jumlah)
	}
	return total
}

func tampilkanDaftarBelanja() {
	fmt.Println("Daftar Belanja:")
	for _, item := range daftarBelanja {
		fmt.Printf("- %s: Rp%.2f x %d = Rp%.2f\n", item.Nama, item.Harga, item.Jumlah, item.Harga*float64(item.Jumlah))
	}
	fmt.Printf("Total: Rp%.2f\n", hitungTotal())
}

func main() {
	tambahItem("Beras", 10000, 2)
	tambahItem("Minyak Goreng", 15000, 1)
	tambahItem("Telur", 2000, 12)
	tambahItem("Gula", -5000, 1)   // validasi harga negatif
	tambahItem("Susu", 12000, 0)   // validasi jumlah nol
	tambahItem("Roti", -10000, -2) // validasi harga dan jumlah negatif
	tampilkanDaftarBelanja()

	for {
		fmt.Println("\n==== Menu ====")
		fmt.Println("1. Tambah Item")
		fmt.Println("2. Tampilkan Daftar Belanja")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih menu (1-3): ")

		var pilihan int
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			var nama string
			var harga float64
			var jumlah int

			fmt.Print("Nama Item: ")
			fmt.Scan(&nama)
			fmt.Print("Harga: ")
			fmt.Scan(&harga)
			fmt.Print("Jumlah: ")
			fmt.Scan(&jumlah)
			tambahItem(nama, harga, jumlah)
		} else if pilihan == 2 {
			tampilkanDaftarBelanja()
		} else if pilihan == 3 {
			fmt.Println("Terima kasih!")
			break
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
