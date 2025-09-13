package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Erro ao ler arquivo", http.StatusBadRequest)
		return
	}
	defer file.Close()

	uploadDir := `C:\Users\negre\OneDrive\Desktop\teste`

	err = os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		http.Error(w, "Erro ao criar pasta de uploads", http.StatusInternalServerError)
		return
	}

	dst, err := os.Create(uploadDir + `\` + header.Filename)
	if err != nil {
		http.Error(w, "Erro ao salvar arquivo", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Erro ao gravar arquivo", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Upload de %s concluído! Arquivo salvo em %s\\%s", header.Filename, uploadDir, header.Filename)
}

func main() {
	os.MkdirAll("uploads", os.ModePerm)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/upload", uploadHandler)

	fmt.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
