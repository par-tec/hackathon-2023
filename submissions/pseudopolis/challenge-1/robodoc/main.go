package main

import (
	"bytes"
	"context"
	"crypto/tls"
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Chat struct {
	Model            string        `json:"model"`
	Messages         []ChatMessage `json:"messages"`
	Temperature      float64       `json:"temperature"`
	MaxTokens        int           `json:"max_tokens"`
	TopP             int           `json:"top_p"`
	FrequencyPenalty int           `json:"frequency_penalty"`
	PresencePenalty  int           `json:"presence_penalty"`
	Stop             []string      `json:"stop"`
}

type ChatResponse struct {
	Choices []struct {
		FinishReason string `json:"finish_reason"`
		Index        int    `json:"index"`
		Message      struct {
			Content string `json:"content"`
			Role    string `json:"role"`
		} `json:"message"`
	} `json:"choices"`
	Created int    `json:"created"`
	ID      string `json:"id"`
	Model   string `json:"model"`
	Object  string `json:"object"`
	Usage   struct {
		CompletionTokens int `json:"completion_tokens"`
		PromptTokens     int `json:"prompt_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

func authenticateRequest(req *http.Request) {

	openAIKey := os.Getenv("OPENAI_KEY")
	if openAIKey == "" {
		log.Fatal("nessuna chiave OPENAI trovata")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+openAIKey)
}

const openAIEndpoint = "https://api.openai.com/v1/chat/completions"

const FILEMESSAGES string = "messages.json"

//go:embed messages.json
var messagesByte []byte

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	// Richiesta utente
	var richiesta ChatMessage
	richiesta.Role = "user"
	richiesta.Content = strings.Join(os.Args[1:], " ")

	/*
		msgContent, err := os.ReadFile(FILEMESSAGES)
		if err != nil {
			log.Fatalf("impossibile leggere file %s: %s", FILEMESSAGES, err.Error())
		}
	*/

	chat := Chat{
		Model: "gpt-4",
		// Messages:         messages,
		Temperature:      0.96,
		MaxTokens:        250,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
		Stop:             []string{"|STOP|"},
	}

	var msgReader = bytes.NewReader(messagesByte)
	var messages []ChatMessage
	if err := json.NewDecoder(msgReader).Decode(&messages); err != nil {
		log.Fatalf("decodifica json dei messaggi in errore: %s", err.Error())
	}

	messages = append(messages, richiesta)

	chat.Messages = messages
	// fmt.Printf("%+v", chat)

	payload, err := json.Marshal(chat)
	if err != nil {
		log.Fatalf("parsing json in errore: %v", err)
	}

	req, err := http.NewRequestWithContext(
		ctx, "POST", openAIEndpoint, bytes.NewBuffer(payload))
	if err != nil {
		log.Fatalf("creazione request in errore: %v", err)
	}

	authenticateRequest(req)

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

	fmt.Println(response)
}
