Conduct conversations
We recommend that you create conversational generation and chatbots by using models that support the chat completion endpoint. The chat completion models and endpoint require a different input structure than the completion endpoint.
The API is adept at carrying on conversations with humans and even with itself. With just a few lines of instruction, the API can perform as a customer service chatbot that intelligently answers questions without getting flustered, or a wise-cracking conversation partner that makes jokes and puns. The key is to tell the API how it should behave and then provide a few examples.
In this demonstration, the API supplies the role of an AI answering questions:
ConsoleCopy
The following is a conversation with an AI assistant. The assistant is helpful, creative, clever, and very friendly.

Human: Hello, who are you?
AI: I am an AI created by OpenAI. How can I help you today?
Human: 
Let's look at a variation for a chatbot named "Cramer," an amusing and somewhat helpful virtual assistant. To help the API understand the character of the role, you provide a few examples of questions and answers. All it takes is just a few sarcastic responses and the API can pick up the pattern and provide an endless number of similar responses.
ConsoleCopy
Cramer is a chatbot that reluctantly answers questions.

###
User: How many pounds are in a kilogram?
Cramer: This again? There are 2.2 pounds in a kilogram. Please make a note of this.
###
User: What does HTML stand for?
Cramer: Was Google too busy? Hypertext Markup Language. The T is for try to ask better questions in the future.
###
User: When did the first airplane fly?
Cramer: On December 17, 1903, Wilbur and Orville Wright made the first flights. I wish they'd come and take me away.
###
User: Who was the first man in space?
Cramer: 
Guidelines for designing conversations
Our demonstrations show how easily you can create a chatbot that's capable of carrying on a conversation. Although it looks simple, this approach follows several important guidelines:
•	Define the intent of the conversation. Just like the other prompts, you describe the intent of the interaction to the API. In this case, "a conversation." This input prepares the API to process subsequent input according to the initial intent.
•	Tell the API how to behave. A key detail in this demonstration is the explicit instructions for how the API should interact: "The assistant is helpful, creative, clever, and very friendly." Without your explicit instructions, the API might stray and mimic the human it's interacting with. The API might become unfriendly or exhibit other undesirable behavior.
•	Give the API an identity. At the start, you have the API respond as an AI created by OpenAI. While the API has no intrinsic identity, the character description helps the API respond in a way that's as close to the truth as possible. You can use character identity descriptions in other ways to create different kinds of chatbots. If you tell the API to respond as a research scientist in biology, you receive intelligent and thoughtful comments from the API similar to what you'd expect from someone with that background.

