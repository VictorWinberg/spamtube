import os
from revChatGPT.V1 import Chatbot

if any(v is None for v in [os.getenv('GOOGLE_EMAIL'), os.getenv('GOOGLE_PASSWORD'), os.getenv('VOICE_INPUT')]):
    print('ChatGPT missing required environment variables')
    quit()

chatbot = Chatbot(config={
    "email": os.getenv('GOOGLE_EMAIL'),
    "password": os.getenv('GOOGLE_PASSWORD')
})

voice_input = os.getenv('VOICE_INPUT')

prompt_title = "Please provide a clickbait title for the following text with a maximum of 42 characters: " + voice_input
prompt_summarize = "Please summarize the following text into 40 seconds, but using the first person point of view: " + voice_input

for data in chatbot.ask(prompt_title):
    title = data["message"]
for data in chatbot.ask(prompt_summarize):
    summarize = data["message"]

with open('out/title.txt', 'w') as f:
    f.write(title.replace('"', ''))
with open('out/text.txt', 'w') as f:
    f.write(summarize)

print(f'Successfully generated a ChatGPT response')
