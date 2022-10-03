package main

import (
	"flag"
	"os"

	htgotts "github.com/hegedustibor/htgo-tts"
	handlers "github.com/hegedustibor/htgo-tts/handlers"
	voices "github.com/hegedustibor/htgo-tts/voices"
)

const OUTPUT_FOLDER = "out"
const OUTPUT_NAME = "audio"

func main() {
	text := flag.String("text", getEnv("VOICE_INPUT", "Hello SpamTube"), "Text-to-Speech input")
	flag.Parse()

	speech := htgotts.Speech{Folder: OUTPUT_FOLDER, Language: voices.Danish, Handler: &handlers.MPlayer{}}
	os.Remove(speech.Folder + "/" + OUTPUT_NAME + ".mp3")
	speech.CreateSpeechFile(*text, OUTPUT_NAME)
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}
