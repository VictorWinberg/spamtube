package generator

import (
	"fmt"
	"log"
	"os"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

const IMAGE_EXT = "png"
const TEST_IMAGE_EXT = "jpg"
const AUDIO_EXT = "mp3"
const OUTPUT_FILE = "./out/video.mp4"
const PREFIX = "scaled_"
const VIDEO_LENGTH = 30 // In seconds
const OUTPUT_VIDEO_FORMAT = "yuvj420p"
const VIDEO_CODEC = "libx264"
const OUTPUT_FPS = 30

func CreateVideo() {
	images, imagePath, imageExt := GetImages()
	_, audioPath := GetAudio()
	ScaleImages(images, imagePath)
	err := GenerateVideo(imagePath, imageExt, audioPath)
	DeleteScaledImages(images, imagePath)
	fmt.Println(err)
}

func ScaleImages(images []os.DirEntry, path string) {
	for _, f := range images {
		fmt.Printf("Found file to scale: %s%s\n", path, f.Name())
		ffmpeg.Input(path+f.Name()).Output(fmt.Sprintf("%s%s%s", path, PREFIX, f.Name()), ffmpeg.KwArgs{"vf": "scale=320:240"}).OverWriteOutput().Run()
	}
}

func GenerateVideo(imagePath string, imageExt string, audioPath string) error {
	imageInput := ffmpeg.Input(imagePath+PREFIX+"%03d."+imageExt, ffmpeg.KwArgs{"loop": 1, "framerate": "1/2"})

	audioInput := ffmpeg.Input(audioPath + "audio." + AUDIO_EXT)

	err := ffmpeg.Concat([]*ffmpeg.Stream{imageInput, audioInput}, ffmpeg.KwArgs{"v": 1, "a": 1}).Output(OUTPUT_FILE, ffmpeg.KwArgs{"r": OUTPUT_FPS, "pix_fmt": OUTPUT_VIDEO_FORMAT, "t": VIDEO_LENGTH, "c:v": VIDEO_CODEC}).OverWriteOutput().ErrorToStdOut().Run()
	return err
}

func DeleteScaledImages(images []os.DirEntry, path string) {
	for _, f := range images {
		os.Remove(path + PREFIX + f.Name())
	}
}

func GetImages() ([]os.DirEntry, string, string) {
	path := "./data/images/"
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
		return ExtractTestImages(testFiles), testPath, TEST_IMAGE_EXT
	}
	return imageFiles, path, IMAGE_EXT
}

func GetAudio() ([]os.DirEntry, string) {
	path := "./data/audio/"
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	audioFiles := ExtractAudio(files)
	if len(audioFiles) == 0 {
		fmt.Println("No audio files found, fetching test audio")
		testPath := "./data/test_data/"
		testFiles, err := os.ReadDir(testPath)
		if err != nil {
			log.Fatal(err)
		}
		return ExtractTestAudio(testFiles), testPath
	}
	return audioFiles, path
}

func ExtractFilesWithExt(files []os.DirEntry, searchFileExt string) []os.DirEntry {
	foundFiles := []os.DirEntry{}
	for _, f := range files {
		fileExt := f.Name()[len(f.Name())-len(searchFileExt) : len(f.Name())]
		if fileExt == searchFileExt {
			foundFiles = append(foundFiles, f)
		}
	}
	return foundFiles
}

func ExtractImages(files []os.DirEntry) []os.DirEntry {
	return ExtractFilesWithExt(files, IMAGE_EXT)
}

func ExtractTestImages(files []os.DirEntry) []os.DirEntry {
	return ExtractFilesWithExt(files, TEST_IMAGE_EXT)
}

func ExtractAudio(files []os.DirEntry) []os.DirEntry {
	return ExtractFilesWithExt(files, AUDIO_EXT)
}

func ExtractTestAudio(files []os.DirEntry) []os.DirEntry {
	return ExtractFilesWithExt(files, AUDIO_EXT)
}
