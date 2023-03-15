package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	handlers "github.com/hegedustibor/htgo-tts/handlers"
	voices "github.com/hegedustibor/htgo-tts/voices"
)

const OUTPUT_FOLDER = "out"
const OUTPUT_NAME = "audio"

type Speech struct {
	Folder   string
	Language string
	Handler  handlers.PlayerInterface
}

func main() {
	bytes, err := ioutil.ReadFile("data/text.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(bytes)

	words := strings.Fields(text)

	rand.Seed(time.Now().UnixNano())

	fmt.Printf("Got TTS-input: %s\n", text)
	in := []string{
		voices.Dutch,
		voices.English,
		voices.EnglishAU,
		voices.EnglishUK,
		voices.Hindi,
		voices.Norwegian,
	}
	language := in[rand.Intn(len(in))]
	fmt.Printf("Language: %s\n", language)
	speech := Speech{Folder: OUTPUT_FOLDER, Language: language, Handler: &handlers.MPlayer{}}
	os.Remove(speech.Folder + "/" + OUTPUT_NAME + ".mp3")

	var input string
	for _, word := range words {
		if len(input)+len(word) > 150 {
			err := speech.download("out/audio.mp3", input)
			if err != nil {
				log.Fatal(err)
			}
			input = word
		} else {
			input += word + " "
		}
	}
	err = speech.download("out/audio.mp3", input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Done! Check %s/%s.mp3\n", speech.Folder, OUTPUT_NAME)
}

func (speech *Speech) download(fileName string, text string) error {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("http://translate.google.com/translate_tts?ie=UTF-8&total=1&idx=0&textlen=32&client=tw-ob&q=%s&tl=%s", url.QueryEscape(text), speech.Language)
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	f.Write(body)

	f.Close()
	return nil
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		value = fallback
	}
	return value
}
