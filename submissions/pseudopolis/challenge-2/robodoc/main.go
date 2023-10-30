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
	"robodoc/src/openai"
	"strings"
	"time"
)

const openAIEndpoint = "https://api.openai.com/v1/chat/completions"

const FILEMESSAGES string = "messages.json"

//go:embed messages.json
var messagesByte []byte

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	// Richiesta utente
	nuovaRichiesta := openai.NewRequest(strings.Join(os.Args[1:], " "))

	chat, err := openai.NewChat(messagesByte)
	if err != nil {
		log.Fatal(err)
	}

	chat.Messages = append(chat.Messages, nuovaRichiesta)

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

	openai.AuthenticateRequest(req)

	// fmt.Printf("%v\n", req)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	c := http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		log.Fatalf("invio richiesta http in errore: %s", err.Error())
	}
	defer resp.Body.Close()

	var result openai.ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatalf("json errore in decodifica risposta: %s", err.Error())
	}

	// log.Printf("%v\n", result)

	response := result.Choices[0].Message.Content

	fmt.Println(response)
}
