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

prompt = "Please summarize the following text into 1 minute, but using the first person point of view: " + voice_input
response = ""

for data in chatbot.ask(prompt):
    response = data["message"]

with open('out/text.txt', 'w') as f:
    f.write(response)

print(f'Succesfully generated a chatgpt response')
