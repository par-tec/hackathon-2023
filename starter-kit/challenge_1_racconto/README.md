# Prova 2: scrittura di un breve racconto creativo che si svolge nel mondo immaginario e coi personaggi immaginari creati nella prova 1 (Durata: 1.5 ore)

# Obiettivo:
    Acquisire familiarità con le API, di OpenAI o altri algoritmi, per la generazione di testo in modo molto semplice. Queste attività consentiranno ai partecipanti di sperimentare e comprendere come utilizzare la generazione di testo in contesti pratici. Una volta che avranno familiarità con questa capacità di base, potranno affrontare sfide più complesse e integrare la tecnologia in applicazioni più avanzate.
    - Descrizione:
        Utilizzare l'API di base per GPT-3.5 (Azure) o altri servizi. Creare un piccolo script che generi frasi, idee o domande casuali utilizzando l'IA. Questa prova ha lo scopo di comprendere come inviare richieste alle API e ottenere una risposta semplice. I testi casuali così creati dovranno essere raccolti e utilizzati per costruire un breve racconto creativo che si svolge nel mondo immaginario e coi personaggi creati nella Challenge di "introduzione e onboarding". Il risultato può essere una storia breve o un articolo. Per creare questo risultato finale dovranno fare attenzione per utilizzare l’API al fine di ottenere testi utili allo scopo ed essere creativi e originali nello scrivere il racconto o l’articolo.

# Istruzioni:
    1.	Per l'accesso alle API OpenAI tramite Azure, fare riferimento agli esempi di codice forniti nello starter-kit.
    2.	Scrivere uno script che invii una richiesta API per generare un breve testo casuale, accettando in input degli argomenti correlati al mondo virtuale della prova 1
    3.	Il testo generato potrebbe essere una breve affermazione, una domanda/risposta o un'idea casuale
    4.	I partecipanti dovrebbero essere in grado di stampare e salvare il testo generato nello script
    5.	Raccogliere queste frasi e costruire un piccolo racconto/articolo creativo di massimo una pagina, salvare il racconto/articolo

# Criteri di valutazione:
    1.	Funzionalità: La soluzione dell'applicazione genera testo in modo efficace e risponde alle richieste degli utenti senza errori evidenti
    2.	Qualità del Testo Generato: La qualità del testo generato è elevata. È ben strutturato, comprensibile e libero da errori grammaticali o semantici
    3.	Facilità d'Uso: L'applicazione è intuitiva e facile da usare per gli utenti
    4.	Tempo di Risposta: L'applicazione genera il testo in un tempo ragionevole. Le risposte sono rapide
    5.	Architettura: Il codice è ben scritto, modulare, quanto è facile aggiungere funzionalità
    6.	Originalità e Creatività: La soluzione proposta mostra originalità e creatività nell'uso delle API per generare testo
    7.	Sicurezza e Conformità: L'applicazione rispetta le linee guida di sicurezza e conformità quando si lavora con le API. Per esempio protezione delle chiavi

# Esempio di chiamata alle API OpenAI:

# PYTHON:
    import os
    import openai
    openai.api_type = "azure"
    openai.api_base = "https://hackarome1.openai.azure.com/"
    openai.api_version = "2023-07-01-preview"
    openai.api_key = os.getenv("OPENAI_API_KEY")

    response = openai.ChatCompletion.create(
    engine="saGPT",
    messages = [{"role":"system","content":"You are an AI assistant that helps people find information."},{"role":"user","content":"Ciao, se a Roma sono le 18:00 puoi dirmi che ore sono sulla Luna?"}],
    temperature=0.7,
    max_tokens=800,
    top_p=0.95,
    frequency_penalty=0,
    presence_penalty=0,
    stop=None)

# CURL:
    curl "https://hackarome1openai.openai.azure.com/openai/deployments/saGPT/chat/completions?api-version=2023-07-01-preview" \
    -H "Content-Type: application/json" \
    -H "api-key: YOUR_API_KEY" \
    -d "{
    \"messages\": [{\"role\":\"system\",\"content\":\"You are an AI assistant that helps people find information.\"},{\"role\":\"user\",\"content\":\"Ciao, se a Roma sono le 18:00 puoi dirmi che ore sono sulla Luna?\"}],
    \"max_tokens\": 800,
    \"temperature\": 0.7,
    \"frequency_penalty\": 0,
    \"presence_penalty\": 0,
    \"top_p\": 0.95,
    \"stop\": null
    }" 
 
# C# :
    // Install the .NET library via NuGet: dotnet add package Azure.AI.OpenAI --version 1.0.0-beta.5 
    using Azure;

    using Azure.AI.OpenAI;

    OpenAIClient client = new OpenAIClient(
        new Uri("https://saopenai.openai.azure.com/"),
        new AzureKeyCredential(Environment.GetEnvironmentVariable("AZURE_OPENAI_API_KEY")));

    // ### If streaming is selected
    Response<StreamingChatCompletions> response = await client.GetChatCompletionsStreamingAsync(
        deploymentOrModelName: "saGPT",
        new ChatCompletionsOptions()
        {
            Messages =
            {
                new ChatMessage(ChatRole.System, @"You are an AI assistant that helps people find information."),
                new ChatMessage(ChatRole.User, @"Ciao, se a Roma sono le 18:00 puoi dirmi che ore sono sulla Luna?"),
            },
            Temperature = (float)0.7,
            MaxTokens = 800,
            NucleusSamplingFactor = (float)0.95,
            FrequencyPenalty = 0,
            PresencePenalty = 0,
        });
    using StreamingChatCompletions streamingChatCompletions = response.Value;


    // ### If streaming is not selected
    Response<ChatCompletions> responseWithoutStream = await client.GetChatCompletionsAsync(
        "saGPT",
        new ChatCompletionsOptions()
        {
            Messages =
            {
                new ChatMessage(ChatRole.System, @"You are an AI assistant that helps people find information."),
                new ChatMessage(ChatRole.User, @"Ciao, se a Roma sono le 18:00 puoi dirmi che ore sono sulla Luna?"),
            },
            Temperature = (float)0.7,
            MaxTokens = 800,
            NucleusSamplingFactor = (float)0.95,
            FrequencyPenalty = 0,
            PresencePenalty = 0,
        });

    ChatCompletions completions = responseWithoutStream.Value;
