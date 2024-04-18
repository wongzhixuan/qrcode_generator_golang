package file

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"qrcode_generator/app/utils/random"
	"time"
)

func SaveImage(imageData image.Image, targetPath string, filePfx string, fileType string) (*os.File, error) {
	filePath := fmt.Sprintf("%s/%s%s%v.%s", targetPath, filePfx, random.RandomString(5), time.Now().UTC().Unix(), fileType)
	f, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	switch fileType {
	case "png":
		err = png.Encode(f, imageData)
		if err != nil {
			return nil, err
		}
	case "jpeg":
		err = jpeg.Encode(f, imageData, nil)
		if err != nil {
			return nil, err
		}
	default:
		err = png.Encode(f, imageData)
		if err != nil {
			return nil, err
		}
	}

	return f, nil
}

func ReadFile(filePath string) (*os.File, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return f, nil
}
