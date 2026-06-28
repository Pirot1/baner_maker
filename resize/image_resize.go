package resize

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
)

func Image_resize(fileName string, width int, height int) image.Image {
	filePath := filepath.Join("images", "input", fileName)
	inFile, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Can't open image %s: %v\n", filePath, err)
		return nil
	}
	defer inFile.Close()

	var img image.Image
	ext := strings.ToLower(filepath.Ext(fileName))
	switch ext {
	case ".jpg", ".jpeg":
		img, err = jpeg.Decode(inFile)
	case ".png":
		img, err = png.Decode(inFile)
	default:
		fmt.Printf("Unsupported format for resize: %s\n", ext)
		return nil
	}

	if err != nil {
		fmt.Printf("Can't decode image %s: %v\n", filePath, err)
		return nil
	}

	resizedImg := imaging.Resize(img, width, height, imaging.Lanczos)
	mirroredImg := imaging.FlipH(resizedImg)
	finalImg := imaging.FlipV(mirroredImg)

	return finalImg
}
