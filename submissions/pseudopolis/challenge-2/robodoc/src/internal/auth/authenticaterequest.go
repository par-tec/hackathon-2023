package auth

import (
	_ "embed"
	"log"
	"net/http"
	"os"
)

//go:embed key.txt
var keyByte []byte

func AuthenticateRequest(req *http.Request) {
	openAIKey := os.Getenv("OPENAI_KEY")
	if openAIKey == "" {
		openAIKey = string(keyByte)
		log.Print("nessuna chiave OPENAI tra le variabili d'ambiente\nuso chiave di test che verrà rimossa in seguito")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+openAIKey)
}
