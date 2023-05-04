import whisper
from datetime import timedelta
import os

model = whisper.load_model("base")

transcribe = model.transcribe("data/audio.mp3")
segments = transcribe['segments']

for segment in segments:
    startTime = str(0)+str(timedelta(seconds=int(segment['start'])))+',000'
    endTime = str(0)+str(timedelta(seconds=int(segment['end'])))+',000'
    text = segment['text']
    segmentId = segment['id']+1
    segment = f"{segmentId}\n{startTime} --> {endTime}\n{text[1:] if text[0] == ' ' else text}\n\n"

    srtFilename = os.path.join("out", "subs.srt")
    with open(srtFilename, 'a', encoding='utf-8') as srtFile:
        srtFile.write(segment)

print(f'Successfully generated subtitles with {len(segments)} lines')