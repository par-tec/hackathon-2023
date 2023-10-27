Trigger ideas
One of the most powerful yet simplest tasks you can accomplish with the API is generating new ideas or versions of input. Suppose you're writing a mystery novel and you need some story ideas. You can give the API a list of a few ideas and it tries to add more ideas to your list. The API can create business plans, character descriptions, marketing slogans, and much more from just a small handful of examples.
In the next demonstration, you use the API to create more examples for how to use virtual reality in the classroom:
ConsoleCopy
Ideas involving education and virtual reality

1. Virtual Mars
Students get to explore Mars via virtual reality and go on missions to collect and catalog what they see.

2.
This demonstration provides the API with a basic description for your list along with one list item. Then you use an incomplete prompt of "2." to trigger a response from the API. The API interprets the incomplete entry as a request to generate similar items and add them to your list.
Guidelines for triggering ideas
Although this demonstration uses a simple prompt, it highlights several guidelines for triggering new ideas:
•	Explain the intent of the list. Similar to the demonstration for the text classifier, you start by telling the API what the list is about. This approach helps the API to focus on completing the list rather than trying to determine patterns by analyzing the text.
•	Set the pattern for the items in the list. When you provide a one-sentence description, the API tries to follow that pattern when generating new items for the list. If you want a more verbose response, you need to establish that intent with more detailed text input to the API.
•	Prompt the API with an incomplete entry to trigger new ideas. When the API encounters text that seems incomplete, such as the prompt text "2.," it first tries to determine any text that might complete the entry. Because the demonstration had a list title and an example with the number "1." and accompanying text, the API interpreted the incomplete prompt text "2." as a request to continue adding items to the list.
•	Explore advanced generation techniques. You can improve the quality of the responses by making a longer more diverse list in your prompt. One approach is to start with one example, let the API generate more examples, and then select the examples you like best and add them to the list. A few more high-quality variations in your examples can dramatically improve the quality of the responses.

