package main

import (
	"fmt"
	"net/http"
	"strings"
	"unicode/utf8"
)

type data_pendaftaran struct {
	nama string
	prodi string
}

var list_pendaftar []data_pendaftaran

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", homeHandler)
	mux.HandleFunc("GET /form", formHandler)
	mux.HandleFunc("POST /simpan", simpanHandler)
	http.ListenAndServe(":8080", mux)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	tag := r.URL.Query().Get("tag")
	fmt.Fprintln(w, "All Query String:", q)
	fmt.Fprintln(w, "First Tag:", tag)
	
	all_tags := q["tag"]
	fmt.Fprintln(w, "All Tags:", strings.Join(all_tags, "#"))
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	html := tampilkanForm(nil,nil)
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}

func simpanHandler(w http.ResponseWriter, r *http.Request) {
	nama := r.PostFormValue("nama")
	prodi := r.PostFormValue("prodi")
	
	errors := []string{}
	if nama == "" {
		errors = append(errors, "Nama harus diisi")
	} else if utf8.RuneCountInString(nama) > 20 {
		errors = append(errors, "Nama maksimal 20 karakter, terlalu panjang boss, ganti nama sana")
	}
	if prodi != "2404130000" && prodi != "2404140000" {
		errors = append(errors, "Prodi tidak valid")
	}
	if len(errors) > 0 {
		values := []string{nama, prodi}
		html := tampilkanForm(values, errors)
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, html)
		return
	}
	list_pendaftar = append(list_pendaftar, data_pendaftaran{nama: nama, prodi: prodi})

	http.Redirect(w, r, "/form", http.StatusSeeOther)
}

func tampilkanForm(values, errors []string) string {
	html := `<h4>Form Pendaftaran</h4>
	<p style="color:red">%s</p>
	<form method="POST" action="/simpan">
		<table>
			<tr>
				<td>Nama:</td>
				<td><input type="text" value="%s" name="nama"></td>
			</tr>
			<tr>
				<td>NIM:</td>
				<td><input type="number" value="%s" name="NIM"></td>
			</tr>
			<tr>
				<td>Prodi:</td>
				<td>
					<select name="prodi">
						<option value="2404130000">Teknik Informatika</option>
						<option value="2404140000">Sistem Informasi</option>
					</select>
				</td>
			</tr>
			<tr>
				<td></td>
				<td><button type="submit">Kirim</button></td>
			</tr>
		</table>
	</form>`
	txt_error := ""
	if errors != nil {
		txt_error = strings.Join(errors, "<br />")
	}
	nama := ""
	if len(values) > 0 {
		nama = values[0]
	}

	html += "<ol>"
	for _, v := range list_pendaftar {
		html += "<li>" + v.nama + " - " + v.prodi + "</li>"
	}
	html += "</ol>"

	return fmt.Sprintf(html, txt_error, nama)
}