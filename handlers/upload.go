package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/negre/dropbox/utils"
)

const UploadDir = `C:\Users\negre\OneDrive\Desktop\teste`

func EnsureUploadDir() {
	if err := os.MkdirAll(UploadDir, os.ModePerm); err != nil {
		panic("Não foi possível criar a pasta de uploads: " + err.Error())
	}
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
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

	if err := utils.SaveFile(UploadDir, header.Filename, file); err != nil {
		http.Error(w, "Erro ao salvar arquivo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Upload de %s concluído! Arquivo salvo em %s", header.Filename, UploadDir)
}
