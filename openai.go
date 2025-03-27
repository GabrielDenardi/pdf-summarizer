package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const openAIEndpoint = "https://api.openai.com/v1/chat/completions"

type OpenAIRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
}

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIResponse struct {
	Choices []struct {
		Message ChatMessage `json:"message"`
	} `json:"choices"`
}

func summarizeText(text string) (string, error) {

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("chave da API OpenAI não definida")
	}

	prompt := fmt.Sprintf("Resuma o seguinte texto de forma clara e concisa:\n\n%s", text)

	reqBody := OpenAIRequest{
		Model: "gpt-4o",
		Messages: []ChatMessage{
			{Role: "system", Content: "Você é um assistente que resume textos."},
			{Role: "user", Content: prompt},
		},
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", openAIEndpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		fmt.Printf("Erro na API OpenAI: status %d, body: %s\n", resp.StatusCode, string(bodyBytes))
		return "", fmt.Errorf("erro na API OpenAI: status %d", resp.StatusCode)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("erro na API OpenAI: status %d", resp.StatusCode)
	}

	var response OpenAIResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	if len(response.Choices) == 0 {
		return "", fmt.Errorf("nenhuma resposta recebida da API")
	}

	return response.Choices[0].Message.Content, nil
}
