package main

import (
	"fmt"
)

var version = "1.0.0" // variabel level package

type Animal struct {
	Name string
	Age  int
	JumlahKaki int
	bisaTerbang bool
	sukaKeju bool
}



func main() {
	// variabel local
	/* var hewan string
	var nilai int
	var ipk float64
	var alamat = "Jakarta"
	var isLulus = true //otomatis akan menggunakan tipe boolean

	nama := "Budi" // variabel short declaration
	bisaTerbang := false */

	// a := 12
	// b := 3.14

	// c := a + b // error karena tipe data berbeda
	// kecuali jika kita melakukan konversi tipe data
	// bisa juga c := float64(a) + b, c akan menjadi float64

	d, e := tambah_kurang(10, 5) // bisa mengembalikan lebih dari satu nilai, maka harus ada tipe data returnnya
	fmt.Println("Hasil penjumlahan:", d)
	fmt.Println("Hasil pengurangan:", e)

	ipk := 2.0
	if ipk >= 3.5 {
		fmt.Println("Selamat, Anda lulus dengan predikat cumlaude!")
	}else if ipk >= 3.0 {
		fmt.Println("Selamat, Anda lulus dengan predikat sangat memuaskan!")
	}else {
		fmt.Println("IPK anda", ipk, "Anda tidak lulus dengan predikat memuaskan.")
	}

	switch ipk { //lebih baik menggunakan switch case daripada if else jika banyak kondisi
	case 4.0:
		fmt.Println("Selamat, Anda lulus dengan predikat cumlaude!")
	case 3.5:
		fmt.Println("Selamat, Anda lulus dengan predikat sangat memuaskan!")
	default:
		fmt.Println("IPK anda", ipk, "Anda tidak lulus dengan predikat memuaskan.")
	}

	//klasik loop
	for i := 0; i < 5; i++ {
		fmt.Printf("Perulangan ke-%d\n", i) //&d adalah format untuk menampilkan angka desimal
	}

	/*
	loop yang lain adalah:
	for {
		// perulangan tanpa kondisi berhenti
		fmt.Println("Perulangan tanpa kondisi berhenti")
	} // ini adalah loop yang tidak ada kondisi berhentinya, maka akan terus berjalan sampai program dihentikan
	*/

	var kucing Animal
	kucing.Name = "Tommy"
	kucing.Age = 2
	kucing.JumlahKaki = 4
	kucing.bisaTerbang = false
	kucing.sukaKeju = false
	fmt.Printf("Nama: %s, Usia: %d, Jumlah Kaki: %d, Bisa Terbang: %t, Suka Keju: %t\n", kucing.Name, kucing.Age, kucing.JumlahKaki, kucing.bisaTerbang, kucing.sukaKeju)

	tikus := Animal{
		Name: "Jerry",
		Age:  1,
		JumlahKaki: 4,
		bisaTerbang: false,
		sukaKeju: true,
	}
	fmt.Printf("Nama: %s, Usia: %d, Jumlah Kaki: %d, Bisa Terbang: %t, Suka Keju: %t\n", tikus.Name, tikus.Age, tikus.JumlahKaki, tikus.bisaTerbang, tikus.sukaKeju)
	// %s dan %t adalah format untuk menampilkan string dan boolean

	//slice adalah tipe data yang bisa menampung banyak data, mirip seperti array, tapi lebih fleksibel
	hewan := []string{"kucing", "tikus", "anjing", "burung"} //slice of string
	fmt.Println("Hewan:", hewan)
	
	nilai := make([]int, 5) //slice of int
	for i := 0; i < len(nilai); i++ {
		nilai[i] = i + 1
	}
	fmt.Println("Nilai:", nilai)
	nilai = append(nilai, 6) // menambahkan data ke slice
	fmt.Println("Nilai setelah ditambahkan:", nilai)
	nilai = append(nilai, 7, 8, 9) // menambahkan banyak data ke slice
	fmt.Println("Nilai setelah ditambahkan banyak data:", nilai)

	nilai2 := make([]int, 0, 5) //slice of int dengan panjang 0 dan kapasitas 5

	nilai2 = append(nilai2, 1, 2, 3, 4, 5) // menambahkan banyak data ke slice

	for i, n := range nilai2 { //range adalah cara untuk melakukan perulangan pada slice
		fmt.Printf("Index: %d, Nilai: %d\n", i, n)
	}
	for _, n := range nilai2 { //jika tidak ingin menggunakan index, bisa menggunakan _ untuk mengabaikan index
		fmt.Printf("Nilai: %d\n", n)
	}

	//ada yang mirip dengan slice, yaitu map, yaitu tipe data yang bisa menampung banyak data dengan key dan value
	//map[string]int adalah map dengan key string dan value int
	//map[string]string adalah map dengan key string dan value string
	//map[int]string adalah map dengan key int dan value string
	//map[int]int adalah map dengan key int dan value int

	kaki_hewan := map[string]int{
		"kucing": 4,
		"tikus": 4,
		"anjing": 4,
		"burung": 2,
	}
	fmt.Println("Jumlah kaki hewan:", kaki_hewan)
	fmt.Println("Jumlah kaki kucing:", kaki_hewan["kucing"])

	//pointer adalah tipe data yang menyimpan alamat memori dari variabel lain
	b := 5
	fmt.Println("Nilai b sebelum diubah:", b)
	ganti_angka(&b)
	fmt.Println("Nilai b setelah diubah:", b)

	// data, err := os.ReadFile("data.txt") //membaca file data.txt
	// if err != nil {
	// 	fmt.Println("Error membaca file:", err)
	// 	return
	// } kadang pakai fmt.Errorf untuk menampilkan error, kadang pakai log.Fatal untuk menampilkan error dan menghentikan program
}

func ganti_angka(a *int) {
	*a = 10 //mengubah nilai dari variabel yang di pointer
}

func hitungLuasPersegiPanjang(panjang int, lebar int) int /*harus ada tipe data returnnya*/{
	return panjang * lebar
}

func tambah_kurang(a int, b int) (int, int) /* bisa mengembalikan lebih dari satu nilai, maka harus ada tipe data returnnya*/{
	return a + b, a - b
}
