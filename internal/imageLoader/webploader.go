package imageloader

import (
	"fmt"
	"image"
	"io"
	"os"

	"golang.org/x/image/webp"
)

func DecodeWebP(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	// Decode the WebP file
	img, err := decodewebp(file)
	if err != nil {
		fmt.Println("Error decoding WebP file:", err)
		return nil, err
	}
	return img, nil
}

func decodewebp(r io.Reader) (image.Image, error) {
	return webp.Decode(r)
}
