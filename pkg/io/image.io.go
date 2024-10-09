package io

import (
	"image"
	"image/jpeg"
	"log"
	"os"
)

func LoadImage(path string) image.Image {
	inputFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	// Step 2: Decode the JPEG image
	img, err := jpeg.Decode(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	return img

}
