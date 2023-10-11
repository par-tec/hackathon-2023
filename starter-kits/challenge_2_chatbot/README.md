# Prova 3: Creazione di un Chatbot in grado di rispondere a domande sul testo creativo ottenuto nella prova 2 (Durata: 1.5 ore)

    •	Obiettivo: Sviluppare un semplice chatbot in grado di rispondere a domande specifiche riguardanti il racconto o l’articolo generato nella Prova 2
    •	Descrizione: Gli scritti creativi ottenuti con la prova 2 saranno distribuiti ai partecipanti, eventualmente utilizzeremo testi già pronti in caso di necessità. Dovranno creare un chatbot che possa rispondere in modo accurato alle domande poste dagli utenti basate sul testo ricevuto. Utilizzeranno le API OpenAI, o di altri, per generare risposte informative e coerenti.

# Istruzioni:
    1.	Utilizzare il testo creativo generato nella prova 2 come input, su questo testo dovranno essere basate le risposte del chatbot.
    2.	Sviluppare un interfaccia che accetti domande dagli utenti.
    3.	Quando un utente inserisce una domanda, questa è inviata al chatbot tramite API OpenAI, il chatbot analizza il testo e genera una risposta basata sul contenuto del racconto/articolo dato in input.
    4.	Il chatbot dev'essere preciso e coerente nel dare le risposte rispetto al contenuto dell'articolo.

# Criteri di Valutazione:
    1. Precisione delle risposte: Il chatbot dovrebbe essere in grado di fornire risposte accurate alle domande inerenti all'articolo. Possibilmente citando i punti salienti dell’articolo che hanno permesso di formulare la risposta
    2. Coerenza: Le risposte dovrebbero essere costruite in modo coerente e leggibile.
    3. Gestione delle domande: Il chatbot dovrebbe essere in grado di interpretare e gestire una varietà di domande poste dagli utenti.
    4. Originalità e Creatività: La soluzione proposta mostra originalità e creatività
    5. L'interfaccia utente aiuta a interagire con il chatbot in modo semplice e intuitivo

    Questa prova consentirà ai partecipanti di applicare la tecnologia in un contesto più specifico e utile, creando un'applicazione pratica di un chatbot che può essere utilizzato per fornire informazioni dettagliate basate sul testo generato.


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
