import os
from decimal import *

input = os.getenv('VOICE_INPUT', 'Hello spamtubers, ...')

words = input.split()

counter = 1
s = float(0)
speed = 9.0

def format_time(s: float) -> str:
    return '{:02}:{:02}:{:02.3f}'.format(int(s//3600), int(s % 3600//60), s % 60)


with open('out/subs.srt' , 'w') as f:
    while (len(words) > 0):
        sentence = words[:8]
        words = words[8:]

        chars = sum([len(w) for w in sentence])

        f.write(str(counter) + '\n' + format_time(s))
        s += chars / speed
        f.write(' --> ' + format_time(s) + '\n' + " ".join(sentence) + '\n\n')

        counter += 1

print(f'Succesfully generated subtitles with {counter} lines')
