package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"

	handlers "github.com/hegedustibor/htgo-tts/handlers"
)

const OUTPUT_FOLDER = "out"
const OUTPUT_NAME = "audio"

type Speech struct {
	Folder   string
	Language string
	Handler  handlers.PlayerInterface
}

type TTSResponse struct {
	Error    int    `json:"Error"`
	Speaker  string `json:"Speaker"`
	Cached   int    `json:"Cached"`
	Text     string `json:"Text"`
	Tasktype string `json:"tasktype"`
	URL      string `json:"URL"`
	Mp3      string `json:"MP3"`
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

	language := "Matthew"
	speech := Speech{Folder: OUTPUT_FOLDER, Language: language, Handler: &handlers.MPlayer{}}
	os.Remove(speech.Folder + "/" + OUTPUT_NAME + ".mp3")

	var input string
	for _, word := range words {
		// do not touch magic number
		if len(input)+len(word) > 150 {
			err := speech.download("out/audio.mp3", input, language)
			if err != nil {
				log.Fatal(err)
			}
			input = word
		} else {
			input += word + " "
		}
	}
	err = speech.download("out/audio.mp3", input, language)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Done! Check %s/%s.mp3\n", speech.Folder, OUTPUT_NAME)
}

func (speech *Speech) download(fileName string, text string, language string) error {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("msg", text)
	_ = writer.WriteField("lang", language)
	_ = writer.WriteField("source", "ttsmp3")
	err = writer.Close()
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://ttsmp3.com/makemp3_new.php", payload)

	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	ttsBody := &TTSResponse{}
	json.NewDecoder(res.Body).Decode(&ttsBody)

	url := fmt.Sprintf("https://ttsmp3.com/dlmp3.php?mp3=%s&location=direct", ttsBody.Mp3)
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
