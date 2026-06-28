package convertor

import (
	"fmt"
	"image"
	"os"
	"path/filepath"
	"strings"

	"github.com/ftrvxmtrx/tga"
)

func Tga_converter(img image.Image, fileName string) {
	ext := filepath.Ext(fileName)
	nameWithoutExt := strings.TrimSuffix(fileName, ext)
	filePath := filepath.Join("images", "output", nameWithoutExt+".tga")
	outputFile, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error while creating TGA file: %v\n", err)
		return
	}
	defer outputFile.Close()

	err = tga.Encode(outputFile, img)
	if err != nil {
		fmt.Printf("Error while saving TGA: %v\n", err)
		return
	}
	fmt.Println("Success!")
}
