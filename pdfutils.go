package main

import (
	"bytes"
	"io"
	"log"

	"github.com/ledongthuc/pdf"
)

func extractTextFromPDF(file io.Reader) (string, error) {
	buf := new(bytes.Buffer)
	_, err := io.Copy(buf, file)
	if err != nil {
		return "", err
	}

	reader := bytes.NewReader(buf.Bytes())
	pdfReader, err := pdf.NewReader(reader, int64(buf.Len()))
	if err != nil {
		return "", err
	}

	var text string
	numPages := pdfReader.NumPage()
	for i := 1; i <= numPages; i++ {
		page := pdfReader.Page(i)
		if page.V.IsNull() {
			continue
		}
		pageText, err := page.GetPlainText(nil)
		if err != nil {
			log.Printf("Erro ao extrair página %d: %v", i, err)
			continue
		}
		text += pageText
	}
	return text, nil
}
