package main

import (
	"fmt"
	"net/http"

	"github.com/negre/dropbox/handlers"
)

func main() {
	handlers.EnsureUploadDir()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/upload", handlers.UploadHandler)

	fmt.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
