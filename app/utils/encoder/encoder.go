package encoder

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

func EncodeFileByFileType(fileType string, f io.Writer, data image.Image) error {
	var err error
	switch fileType {
	case "png":
		err = png.Encode(f, data)
		if err != nil {
			return err
		}
	case "jpeg":
		err = jpeg.Encode(f, data, nil)
		if err != nil {
			return err
		}
	default:
		err = png.Encode(f, data)
		if err != nil {
			return err
		}
	}
	return nil
}
