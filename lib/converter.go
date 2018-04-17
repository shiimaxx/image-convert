package lib

import (
	"image"
	_ "image/gif"
	"image/jpeg"
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
	case "jpeg", "jpg":
		err := convertToJPEG(filename)
		if err != nil {
			return err
		}
	}
	return nil
}

func convertToJPEG(filename string) error {
	srcFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	img, _, err := image.Decode(srcFile)
	if err != nil {
		return err
	}

	destFile, err := os.Create(replaceExt(filename, ".jpg"))
	if err != nil {
		return err
	}
	defer destFile.Close()

	err = jpeg.Encode(destFile, img, nil)
	if err != nil {
		return err
	}
	return nil
}

func convertToPNG(filename string) error {
	srcFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	img, _, err := image.Decode(srcFile)
	if err != nil {
		return err
	}

	destFile, err := os.Create(replaceExt(filename, ".png"))
	if err != nil {
		return err
	}
	defer destFile.Close()

	err = png.Encode(destFile, img)
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
