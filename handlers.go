package main

import (
	"fmt"
	"net/http"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, _, err := r.FormFile("pdf")
	if err != nil {
		http.Error(w, "Erro ao receber o arquivo", http.StatusBadRequest)
		return
	}
	defer file.Close()

	text, err := extractTextFromPDF(file)
	if err != nil {
		http.Error(w, "Erro ao processar o PDF", http.StatusInternalServerError)
		return
	}

	summary, err := summarizeText(text)
	if err != nil {
		http.Error(w, "Erro ao gerar o resumo", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "Resumo:\n%s", summary)
}
