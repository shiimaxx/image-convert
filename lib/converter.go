package converter

import (
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
)

type Image struct {
	Name string
	Type string
}

func convertToPNG(filePath string) error {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return err
	}

	img, _, err := image.Decode(file)
	out, err := os.Create("")
	if err != nil {
		return err
	}

	err = jpeg.Encode(out, img, nil)
	if err != nil {
		return err
	}
	return nil
}

func (i *Image) Convert(destType string) error {
	switch destType {
	case "png":
		err := convertToPNG(i.Path)
		if err != nil {
			return err
		}
	}
	return nil
}

// SearchImageFiles is a
func SearchImageFiles(dir string) ([]string, error) {
	imageFiles := []string{}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		imageFiles = append(imageFiles, path)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return imageFiles, nil
}
