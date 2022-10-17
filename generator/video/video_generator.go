package generator

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strconv"
	"strings"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

const IMAGE_EXT = "png"
const TEST_IMAGE_EXT = "jpg"
const AUDIO_EXT = "mp3"
const OUTPUT_VIDEO_FILE = "./out/video.mp4"
const OUTPUT_AUDIO_FILE = "./out/audio.mp3"
const VIDEO_LENGTH = 30 // In seconds
const OUTPUT_VIDEO_FORMAT = "yuvj420p"
const VIDEO_CODEC = "libx264"
const OUTPUT_FPS = 30
const MIN_AUDIO_FILES = 2
const VOLUME_CHANGE_FACTOR = "0.5"

func CreateVideo() {
	imagePath, imageExt := GetImages()
	audio := GenerateAudio()
	GenerateVideo(imagePath, imageExt, audio)

}

func GenerateVideo(imagePath string, imageExt string, audioInput *ffmpeg.Stream) {
	imageInput := ffmpeg.Input(imagePath+"%03d."+imageExt, ffmpeg.KwArgs{"loop": 1, "framerate": "1/2"})
	err := ffmpeg.Concat([]*ffmpeg.Stream{imageInput, audioInput}, ffmpeg.KwArgs{"v": 1, "a": 1}).Output(OUTPUT_VIDEO_FILE, ffmpeg.KwArgs{"r": OUTPUT_FPS, "pix_fmt": OUTPUT_VIDEO_FORMAT, "t": VIDEO_LENGTH, "c:v": VIDEO_CODEC}).OverWriteOutput().ErrorToStdOut().Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully generated video: " + OUTPUT_VIDEO_FILE)
}

func GenerateAudio() *ffmpeg.Stream {
	pathFiles := GetAudio()
	nFoundFiles := len(Values(pathFiles))
	if nFoundFiles < MIN_AUDIO_FILES {
		log.Fatalf("Expected at least %d audiofiles, found %d", MIN_AUDIO_FILES, nFoundFiles)
	}
	audioInputs := GetAsFfmpegAudioInputs(pathFiles)
	err := ffmpeg.Filter(audioInputs, "amix", ffmpeg.Args{fmt.Sprintf("inputs=%d", nFoundFiles), "duration=longest"}).Output(OUTPUT_AUDIO_FILE, ffmpeg.KwArgs{"c:a": "libmp3lame", "t": 60}).OverWriteOutput().ErrorToStdOut().Run()
	if err != nil {
		fmt.Println("Failed to generate audio for video")
		log.Fatal(err)
	}
	return ffmpeg.Input(OUTPUT_AUDIO_FILE)
}

func GetImages() (string, string) {
	path := "./data/images/"
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	imageFiles := ExtractImages(files)
	if len(imageFiles) == 0 {
		fmt.Println("No images found, fetching test images")
		testPath := "./data/test_data/"
		_, err := os.ReadDir(testPath)
		if err != nil {
			log.Fatal(err)
		}
		return testPath, TEST_IMAGE_EXT
	}
	return path, IMAGE_EXT
}

func GetAudio() map[string][]os.DirEntry {
	paths := []string{"./data/audio/voice/", "./data/audio/music/"}
	pathFiles := map[string][]os.DirEntry{}
	for _, path := range paths {
		files, err := os.ReadDir(path)
		if err != nil {
			continue
		}
		audioFiles := ExtractAudio(files)
		pathFiles[path] = audioFiles
	}
	if len(Values(pathFiles)) == 0 {
		fmt.Println("No audio files found, fetching test audio")
		testPaths := []string{"./data/test_data/"}
		pathFiles = map[string][]os.DirEntry{}
		for _, path := range testPaths {
			testFiles, err := os.ReadDir(path)
			if err != nil {
				log.Fatal(err)
			}
			audioFiles := ExtractTestAudio(testFiles)
			pathFiles[path] = audioFiles
		}
	}
	return pathFiles
}

func GetAsFfmpegAudioInputs(pathFiles map[string][]os.DirEntry) []*ffmpeg.Stream {
	audioInputs := []*ffmpeg.Stream{}
	for path, audioFiles := range pathFiles {
		for _, audioFile := range audioFiles {
			if strings.Contains(path, "music") {
				reducedAudioFile := ReduceVolume(path, audioFile)
				audioInputs = append(audioInputs, reducedAudioFile)
			} else {
				audioInputs = append(audioInputs, ffmpeg.Input(path+audioFile.Name()))
			}
			fmt.Println("Found audio " + audioFile.Name())
		}
	}
	return audioInputs
}

func ReduceVolume(path string, audioFile fs.DirEntry) *ffmpeg.Stream {
	fmt.Println("Trying to reduce volume of file: " + path + audioFile.Name())
	outPutFile := "out/" + audioFile.Name()
	err := ffmpeg.Input(path+audioFile.Name()).Filter("volume", ffmpeg.Args{VOLUME_CHANGE_FACTOR}).Output(outPutFile).OverWriteOutput().Run()
	if err != nil {
		log.Fatal(err)
	}
	d, _ := strconv.ParseFloat(VOLUME_CHANGE_FACTOR, 32)
	fmt.Printf("Reduced volume to %d%% of file: %s as: %s\n", int(d*100), path+audioFile.Name(), outPutFile)
	return ffmpeg.Input(outPutFile)
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

func Values(m map[string][]os.DirEntry) []os.DirEntry {
	r := []os.DirEntry{}
	for _, v := range m {
		r = append(r, v...)
	}
	return r
}
