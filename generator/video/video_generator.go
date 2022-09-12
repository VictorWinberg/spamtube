package generator

import (
	"fmt"
	"log"
	"os"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

const IMAGE_EXT = "jpg"
const OUTPUT_FILE = "./dist/out.mp4"
const PREFIX = "scaled_"
const VIDEO_LENGTH = 30 // In seconds
const OUTPUT_VIDEO_FORMAT = "yuvj420p"
const VIDEO_CODEC = "libx264"
const OUTPUT_FPS = 30

func CreateVideo() {
	images, path := GetImages()
	for _, f := range images {
		fmt.Printf("Found file to scale: %s%s\n", path, f.Name())
		ffmpeg.Input(path+f.Name()).Output(fmt.Sprintf("%s%s%s", path, PREFIX, f.Name()), ffmpeg.KwArgs{"vf": "scale=320:240"}).OverWriteOutput().Run()
	}
	imageInput := ffmpeg.Input(path+PREFIX+"%03d.jpg", ffmpeg.KwArgs{"loop": 1, "framerate": "1/2"})
	audioInput := ffmpeg.Input("./data/test_data/audio.mp3")
	err := ffmpeg.Concat([]*ffmpeg.Stream{imageInput, audioInput}, ffmpeg.KwArgs{"v": 1, "a": 1}).Output(OUTPUT_FILE, ffmpeg.KwArgs{"r": OUTPUT_FPS, "pix_fmt": OUTPUT_VIDEO_FORMAT, "t": VIDEO_LENGTH, "c:v": VIDEO_CODEC}).OverWriteOutput().ErrorToStdOut().Run()
	for _, f := range images {
		os.Remove(path + PREFIX + f.Name())
	}
	fmt.Println(err)
}

func GetImages() ([]os.DirEntry, string) {
	path := "./data/"
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	imageFiles := ExtractImages(files)
	if len(imageFiles) == 0 {
		fmt.Println("No images found, fetching test images")
		testPath := "./data/test_data/"
		testFiles, err := os.ReadDir(testPath)
		if err != nil {
			log.Fatal(err)
		}
		return ExtractImages(testFiles), testPath
	}
	return imageFiles, path
}

func ExtractImages(files []os.DirEntry) []os.DirEntry {
	imageFiles := []os.DirEntry{}
	for _, f := range files {
		fileExt := f.Name()[len(f.Name())-len(IMAGE_EXT) : len(f.Name())]
		if fileExt == IMAGE_EXT {
			imageFiles = append(imageFiles, f)
		}
	}
	return imageFiles
}
