# ROBODOC

Robodoc è un personaggio di fantasia nella città sommersa di Pseudopolis.
E' un robot medico russo che avrebbe bisogno di essere aggiornato ma è ancora capacissimo di stupire con le sue trovate mediche geniali nonstante i rumori che emette fanno pensare che possa smettere di funzionare in ogni momento.

## Codice
Il codice per interfacciarsi con le API di OPENAI è scritto in golang.

### main.go

Nella funzione principale main 
1) si recupera il profilo del personaggio dal file robodoc.json e si embedda come bytes nell'eseguibile,
2) Si crea un assistentevirtuale invocando il package openai
3) si recupera l'interazione utente da passare all'API
4) si stampa la risposta a video la risposta del chatbot

```GO
package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"
	"os"
	"robodoc/src/openai"
	"strings"
	"time"
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

```

### Package OPENAI

Il package OPENAI si occupa di parsare e inviare le richieste all'endpoint OPENAPI

#### Autenticazione
L'autenticazione avviene aggiungendo degli heder alla richiesta http.
Per evitare che la chiave API possa essere usata da altri moduli è inserita all'internto della directory 'internal' nel package auth.


### Compilazione
L'eseguibile robodoc può essere compilato lanciando:

```sh
go build
```

da dentro la directory robodoc.

### Dockerizzazione
Bisogna modificare il file Dockerfile per aggiungere la chiave API da usare. E poi si può creare l'immagine docker lanciando:

```sh
podman build --tag me/robodoc .
```

### Avvio container
```sh
podman run --rm me/robodoc <testo utente>
```
## Video

### Esempio di consulto medico con Robodoc
[![asciicast](https://asciinema.org/a/3uo5aDrd8J6xuDzeFfYE20mqk.svg)](https://asciinema.org/a/3uo5aDrd8J6xuDzeFfYE20mqk)

### Secondo consulto medico
[![asciicast](https://asciinema.org/a/Lvt51GTKpPH5ElvSbsGjf61vZ.svg)](https://asciinema.org/a/Lvt51GTKpPH5ElvSbsGjf61vZ)