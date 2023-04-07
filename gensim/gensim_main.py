import os
from gensim.summarization.summarizer import summarize

with open('out/title.txt', 'w') as f:
    f.write(os.getenv('TITLE'))
with open('out/text.txt', 'w') as f:
    text = os.getenv('VOICE_INPUT')
    summary = summarize(text, word_count = 150)
    f.write(summary)
print(f'Successfully generated a gensim response')
