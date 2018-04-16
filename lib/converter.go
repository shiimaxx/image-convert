package converter

import (
	"image"
	"image/png"
	"os"
	"path/filepath"
)

type Image struct {
	filename string
}

func replaceExt(filePath, newExt string) string {
	ext := filepath.Ext(filePath)
	return filePath[0:len(filePath)-len(ext)] + newExt
}

func (i *Image) ConvertToPNG() error {
	file, err := os.Open(i.filename)
	defer file.Close()
	if err != nil {
		return err
	}

	img, _, err := image.Decode(file)
	out, err := os.Create(replaceExt(i.filename, ".png"))
	if err != nil {
		return err
	}

	err = png.Encode(out, img)
	if err != nil {
		return err
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
