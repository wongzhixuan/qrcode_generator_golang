package converter

import (
	"encoding/base64"
	"fmt"
)

func ConvertImageToHTMLString(imageByte []byte) (string, error) {
	imgBase64 := base64.StdEncoding.EncodeToString(imageByte)
	imgURL := fmt.Sprintf(`data:image/png;base64, %s`, imgBase64)
	imgHTML := fmt.Sprintf(`<img src="%s"/>`, imgURL)
	return imgHTML, nil
}
