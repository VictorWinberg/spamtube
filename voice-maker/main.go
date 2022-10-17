package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	htgotts "github.com/hegedustibor/htgo-tts"
	handlers "github.com/hegedustibor/htgo-tts/handlers"
	voices "github.com/hegedustibor/htgo-tts/voices"
)

const OUTPUT_FOLDER = "out"
const OUTPUT_NAME = "audio"

func main() {
	text := flag.String("text", getEnv("VOICE_INPUT", strings.Repeat("Hello SpamTube ", 10)), "Text-to-Speech input")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	fmt.Printf("Got TTS-input: %s\n", *text)
	in := []string{
		voices.Danish,
		voices.Swedish,
		voices.Norwegian,
		voices.French,
		voices.German,
		voices.Ukrainian,
		voices.EnglishAU,
		voices.EnglishUK,
		voices.Spanish,
	}
	language := in[rand.Intn(len(in))]
	fmt.Printf("Language: %s\n", language)
	speech := htgotts.Speech{Folder: OUTPUT_FOLDER, Language: language, Handler: &handlers.MPlayer{}}
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
