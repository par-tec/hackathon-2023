package openai

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/albertobregliano/hackathon-2023/submissions/pseudopolis/challenge-2/robodoc/src/internal/auth"
)

const openAIEndpoint = "https://api.openai.com/v1/chat/completions"

type AssistantProfile []byte
type AssistenteVirtuale Chat

func NewRequest(s string) ChatMessage {
	// Richiesta utente
	var richiesta ChatMessage
	richiesta.Role = "user"
	richiesta.Content = s
	return richiesta
}

// NuovoAssitenteVirtuale crea un assistente virtuale su GPT basato
// sul profilo passato come argomento.
func NuovoAssitenteVirtuale(assistantData []byte) (AssistenteVirtuale, error) {
	var msgReader = bytes.NewReader(assistantData)
	var messages []ChatMessage
	if err := json.NewDecoder(msgReader).Decode(&messages); err != nil {
		return AssistenteVirtuale{}, fmt.Errorf("decodifica json dei messaggi in errore: %s", err.Error())
	}

	c := Chat{
		Model:            "gpt-4",
		Messages:         messages,
		Temperature:      0.96,
		MaxTokens:        250,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
		Stop:             []string{"|STOP|"},
	}

	return AssistenteVirtuale(c), nil
}

// Chiedi esegue una richiesta di tipo chat all'assistente virtuale.
func (av AssistenteVirtuale) Chiedi(ctx context.Context, domanda string) string {
	req := NewRequest(domanda)
	var assistente = av
	assistente.Messages = append(assistente.Messages, req)
	risposta, err := getResponse(ctx, Chat(assistente))
	if err != nil {
		log.Println(err)
	}
	return risposta
}

func getResponse(ctx context.Context, assistant Chat) (string, error) {
	payload, err := json.Marshal(assistant)
	if err != nil {
		log.Fatalf("parsing json in errore: %v", err)
	}

	req, err := http.NewRequestWithContext(
		ctx, "POST", openAIEndpoint, bytes.NewBuffer(payload))
	if err != nil {
		log.Fatalf("creazione request in errore: %v", err)
	}

	auth.AuthenticateRequest(req)

	// fmt.Printf("%v\n", req)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	c := http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		log.Fatalf("invio richiesta http in errore: %s", err.Error())
	}
	defer resp.Body.Close()

	var result ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatalf("json errore in decodifica risposta: %s", err.Error())
	}

	// log.Printf("%v\n", result)

	response := result.Choices[0].Message.Content

	return response, nil
}
