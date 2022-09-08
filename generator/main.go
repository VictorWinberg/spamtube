package main

import (
	"flag"
	"fmt"
	gen "spamtube/generator/video"
)

var (
	imageFolder = flag.String("folder", "test_data", "The folder where images should be")
)

func main() {
	flag.Parse()
	fmt.Println("Welcome to init() function " + *imageFolder)
	gen.CreateVideo()
}
