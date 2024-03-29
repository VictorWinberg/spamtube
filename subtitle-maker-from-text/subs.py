import os
from decimal import *

f = open("data/text.txt", "r")
input = f.read()

# words = input.split()
words = ["Welcome to Spamtube!"]

counter = 1
s = float(0)
speed = 10.2

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

print(f'Successfully generated subtitles with {counter} lines')
