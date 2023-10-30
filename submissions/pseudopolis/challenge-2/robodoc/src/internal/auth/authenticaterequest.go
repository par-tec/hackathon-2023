package auth

import (
	_ "embed"
	"log"
	"net/http"
	"os"
)

const openAIEndpoint = "https://api.openai.com/v1/chat/completions"

//go:embed key.txt
var keyByte []byte

func AuthenticateRequest(req *http.Request) {
	openAIKey := os.Getenv("OPENAI_KEY")
	if openAIKey == "" {
		openAIKey = string(keyByte)
		log.Print("nessuna chiave OPENAI trovata")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+openAIKey)
}
