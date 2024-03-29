import os
from revChatGPT.V1 import Chatbot

try:
    if any(v is None for v in [os.getenv('GOOGLE_EMAIL'), os.getenv('GOOGLE_PASSWORD'), os.getenv('TEXT_CONTENT')]):
        print('ChatGPT missing required environment variables')
        raise Exception('ChatGPT missing required environment variables')

    chatbot = Chatbot(config={
        "email": os.getenv('GOOGLE_EMAIL'),
        "password": os.getenv('GOOGLE_PASSWORD')
    })

    text_content = os.getenv('TEXT_CONTENT')

    prompt_title = "Please provide a clickbait and dramatic title for the following text with a maximum of 42 characters: " + text_content
    prompt_summarize = "Please summarize the following text into 50 seconds, make it exciting, using the first person point of view, meaning using the pronouns I, me, we, and us, in order to tell a story from the narrator's perspective, censor swear words with symbols, spell out numbers with words: " + text_content

    for data in chatbot.ask(prompt_title):
        title = data["message"]
    for data in chatbot.ask(prompt_summarize):
        summary = data["message"]

    with open('out/title.txt', 'w') as f:
        f.write(title.replace('"', ''))
    with open('out/text.txt', 'w') as f:
        f.write(summary)

    print(f'Successfully generated a ChatGPT response')
except Exception as err: 
    print(err)

    # ChatGPT fallback
    with open('out/title.txt', 'w') as f:
        f.write(os.getenv('TITLE'))

    with open('out/text.txt', 'w') as f:
        text = os.getenv('TEXT_CONTENT')
        f.write(text)
