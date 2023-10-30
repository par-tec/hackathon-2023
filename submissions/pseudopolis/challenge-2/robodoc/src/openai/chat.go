package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type AssistantProfile []byte

func NewAssistant(profileData []byte) []ChatMessage {
	var msgReader = bytes.NewReader(profileData)
	var messages []ChatMessage
	if err := json.NewDecoder(msgReader).Decode(&messages); err != nil {
		log.Fatalf("decodifica json dei messaggi in errore: %s", err.Error())
	}
	return messages
}

func NewRequest(s string) ChatMessage {
	// Richiesta utente
	var richiesta ChatMessage
	richiesta.Role = "user"
	richiesta.Content = s
	return richiesta
}

func NewChat(assistantData []byte) (Chat, error) {
	var msgReader = bytes.NewReader(assistantData)
	var messages []ChatMessage
	if err := json.NewDecoder(msgReader).Decode(&messages); err != nil {
		return Chat{}, fmt.Errorf("decodifica json dei messaggi in errore: %s", err.Error())
	}

	return Chat{
		Model: "gpt-4",
		// Messages:         messages,
		Temperature:      0.96,
		MaxTokens:        250,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
		Stop:             []string{"|STOP|"},
	}, nil
}
