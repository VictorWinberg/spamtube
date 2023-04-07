#!/bin/sh

AUDIO_FILE=data/audio.mp3
OUTPUT_FILE=out/subs.srt

autosrt $AUDIO_FILE -o $OUTPUT_FILE