package handler

import (
	"context"
	"fmt"
	"github.com/labstack/echo"
	"github.com/lil-shimon/workManager/repository"
	"image"
	"log"
	"os"

	vision "cloud.google.com/go/vision/apiv1"
)

// FileReader
//
///**
func FileReader(c echo.Context) (image.Image, error) {
	// first things first, check the data of upload file
	file, err := c.FormFile("file")
	if err != nil {
		return file, err
	}
	fileData := repository.GetImageData(file)
	// then upload to s3

	// get the path of uploaded file
	path := os.Args[0]

	// get text data by file path
	Ocr(path)

	// return text data from upload file
	return fileData, nil
}

// Ocr
// param path string
// return read text
///**
func Ocr(path string) {
	// create context
	ctx := context.Background()

	// Creates a client.
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	defer file.Close()
	image, err := vision.NewImageFromReader(file)
	if err != nil {
		log.Fatalf("Failed to create image: %v", err)
	}

	texts, err := client.DetectTexts(ctx, image, nil, 10)
	if err != nil {
		log.Fatalf("Failed to detect labels: %v", path)
	}

	for _, text := range texts {
		fmt.Println(text.Description)
	}
}
