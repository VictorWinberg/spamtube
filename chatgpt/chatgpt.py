import os
from revChatGPT.V1 import Chatbot
from gensim.summarization.summarizer import summarize

if any(v is None for v in [os.getenv('GOOGLE_EMAIL'), os.getenv('GOOGLE_PASSWORD'), os.getenv('VOICE_INPUT')]):
    print('ChatGPT missing required environment variables')
    quit()

try:
    chatbot = Chatbot(config={
        "email": os.getenv('GOOGLE_EMAIL'),
        "password": os.getenv('GOOGLE_PASSWORD')
    })

    voice_input = os.getenv('VOICE_INPUT')

    prompt_title = "Please provide a clickbait and dramatic title for the following text with a maximum of 42 characters: " + voice_input
    prompt_summarize = "Please summarize the following text into 50 seconds, make it exciting, using the first person point of view, meaning using the pronouns I, me, we, and us, in order to tell a story from the narrator's perspective, also censor swear words with symbols: " + voice_input

    for data in chatbot.ask(prompt_title):
        title = data["message"]
    for data in chatbot.ask(prompt_summarize):
        summary = data["message"]

    with open('out/title.txt', 'w') as f:
        f.write(title.replace('"', ''))
    with open('out/text.txt', 'w') as f:
        f.write(summary)

    print(f'Successfully generated a ChatGPT response')
except:
    with open('out/title.txt', 'w') as f:
        f.write(os.getenv('FALLBACK_TITLE'))
    with open('out/text.txt', 'w') as f:
        text = os.getenv('VOICE_INPUT')
        summary = summarize(text, word_count = 150)
        f.write(summary)
    print(f'Successfully generated a gensim response')
