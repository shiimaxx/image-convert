package lib

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

type Image struct {
	fileName string
}

func replaceExt(filePath, newExt string) string {
	ext := filepath.Ext(filePath)
	return strings.TrimSuffix(filePath, ext) + newExt
}

func Convert(filename, destExt string) error {
	switch destExt {
	case "png":
		err := convertToPNG(filename)
		if err != nil {
			return err
		}
	}
	return nil
}

func convertToPNG(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	out, err := os.Create(replaceExt(fileName, ".png"))
	defer out.Close()
	if err != nil {
		return err
	}

	err = png.Encode(out, img)
	if err != nil {
		return err
	}
	return nil
}

func MakeImageFiles(dir, srcExt string) ([]string, error) {
	nameSuffix := "." + srcExt
	imageFiles := []string{}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == nameSuffix {
			imageFiles = append(imageFiles, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return imageFiles, nil
}
