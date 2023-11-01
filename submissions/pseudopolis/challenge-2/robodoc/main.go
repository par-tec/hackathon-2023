package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/albertobregliano/hackathon-2023/submissions/pseudopolis/challenge-2/robodoc/src/openai"
)

//go:embed robodoc.json
var profiloAssistenteVirtuale []byte

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	// Scelta/creazione assistente viruale.
	robodoc, err := openai.NuovoAssitenteVirtuale(profiloAssistenteVirtuale)
	if err != nil {
		log.Fatal(err)
	}

	// Richiesta utente.
	nuovaRichiesta := strings.Join(os.Args[1:], " ")

	// Risposta dell'assistente viruale.
	risposta := robodoc.Chiedi(ctx, nuovaRichiesta)

	fmt.Println(risposta)
}
