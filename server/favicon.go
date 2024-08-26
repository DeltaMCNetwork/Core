package server

import (
	"encoding/base64"
	"fmt"
	"image"
	_ "image/png"
	"io"
	"os"
)

func loadIcon() string {
	imgPath := "resources/icon.png"
	imgFile, err := os.Open(imgPath)
	if err != nil {
		// file not found or inaccessible
		Warn("Icon not found or inaccessible. Skipping...")
		return ""
	}

	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)
	if err != nil {
		panic("Failed to decode icon: " + err.Error())
	}

	if img.Bounds().Dx()%64 != 0 || img.Bounds().Dy()%64 != 0 {
		// image is not in 64x64 ratio
		panic(fmt.Sprint("Icon should be 64 x 64 pixels. Your icon size: ", img.Bounds().Dx(), " x ", img.Bounds().Dy(), " pixels!"))
	}

	// reset pointer
	imgFile.Seek(0, 0)

	imgData, err := io.ReadAll(imgFile)
	if err != nil {
		panic("Error reading icon: " + err.Error())
	}

	Info("Server favicon loaded!")

	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(imgData)
}
