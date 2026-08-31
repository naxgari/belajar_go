package main

import (
	"fmt"
	"net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	fmt.Println("Server berjalan di port 8080")
	fmt.Println("Tekan Ctrl+c untuk stop")
    
	http.ListenAndServe(":8080", nil)
}