package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	htgotts "github.com/hegedustibor/htgo-tts"
	handlers "github.com/hegedustibor/htgo-tts/handlers"
	voices "github.com/hegedustibor/htgo-tts/voices"
)

const OUTPUT_FOLDER = "out"
const OUTPUT_NAME = "audio"

func main() {
	text := flag.String("text", getEnv("VOICE_INPUT", strings.Repeat("Hello SpamTube ", 10)), "Text-to-Speech input")
	flag.Parse()

	fmt.Printf("Got TTS-input: %s\n", *text)
	speech := htgotts.Speech{Folder: OUTPUT_FOLDER, Language: voices.Swedish, Handler: &handlers.MPlayer{}}
	os.Remove(speech.Folder + "/" + OUTPUT_NAME + ".mp3")
	speech.CreateSpeechFile(*text, OUTPUT_NAME)

	fmt.Printf("Done! Check %s/%s.mp3\n", speech.Folder, OUTPUT_NAME)
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		value = fallback
	}
	return value
}
