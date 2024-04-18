package generator

import (
	"image"
	"image/draw"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/nfnt/resize"
)

type QRCodeData struct {
	Content      string
	Size         int
	HasWatermark bool
	WatermarkImg image.Image
}

// func (code *QRCodeData) GenerateWithWatermarks() (string, error) {

// 	qrImg, err := code.generateQRCode()
// 	if err != nil {
// 		return "", err
// 	}

// 	if code.IsSaveToFile {
// 		f, err := os.Create(fmt.Sprintf("resources/qr_%v.%s", time.Now().UTC().Unix(), code.FileType))
// 		if err != nil {
// 			return "", err
// 		}
// 		defer f.Close()

// 		err = png.Encode(f, m)
// 		if err != nil {
// 			return "", err
// 		}
// 		return f.Name(), nil
// 	}

// 	// Convert QR code image to base64 encoded string
// 	buf := new(bytes.Buffer)
// 	err = png.Encode(buf, m)
// 	if err != nil {
// 		return "", err
// 	}
// 	imgBase64 := base64.StdEncoding.EncodeToString(buf.Bytes())
// 	imgURL := fmt.Sprintf(`data:image/png;base64, %s`, imgBase64)
// 	imgHTML := fmt.Sprintf(`<img src="%s"/>`, imgURL)
// 	return imgHTML, nil
// }

func (code *QRCodeData) Generate() (image.Image, error) {
	qrImg, err := code.generateQRCode()
	if err != nil {
		return nil, err
	}

	if code.HasWatermark {
		if code.WatermarkImg == nil {
			//set default watermark
			f, err := os.Open("assets/default.png")
			if err != nil {
				return nil, err
			}
			defer f.Close()
			readerImg, err := png.Decode(f)
			if err != nil {
				return nil, err
			}
			code.WatermarkImg = readerImg
		}
		qrWithWatermarkImg, err := addWatermarkToQRCode(qrImg, code.WatermarkImg)
		if err != nil {
			return nil, err
		}
		return qrWithWatermarkImg, nil
	}

	return qrImg, nil

}

func (code *QRCodeData) generateQRCode() (image.Image, error) {
	b, err := qr.Encode(code.Content, qr.M, qr.Auto)
	if err != nil {
		return nil, err
	}

	b, err = barcode.Scale(b, code.Size, code.Size)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func resizeWatermark(watermark image.Image, width uint) (image.Image, error) {

	m := resize.Resize(width, 0, watermark, resize.Lanczos3)

	return m, nil
}

func addWatermarkToQRCode(qrCodeImg image.Image, watermarkImg image.Image) (image.Image, error) {
	watermarkWidth := uint(float64(qrCodeImg.Bounds().Dx()) * 0.25)

	watermark, err := resizeWatermark(watermarkImg, watermarkWidth)
	if err != nil {
		return nil, err
	}

	var halfQrCodeWidth, halfWatermarkWidth int = qrCodeImg.Bounds().Dx() / 2, watermark.Bounds().Dx() / 2
	offset := image.Pt(
		halfQrCodeWidth-halfWatermarkWidth,
		halfQrCodeWidth-halfWatermarkWidth,
	)

	watermarkImageBounds := qrCodeImg.Bounds()
	m := image.NewRGBA(watermarkImageBounds)

	draw.Draw(m, watermarkImageBounds, qrCodeImg, image.Point{}, draw.Src)
	draw.Draw(
		m,
		watermark.Bounds().Add(offset),
		watermark,
		image.Point{},
		draw.Over,
	)
	return m, nil
}

// if code.IsSaveToFile{
// 	filePath := fmt.Sprintf("%s", "resources/qrcode")
// 	f, err := file.SaveImage(qrImg,filePath, "qrcode", code.FileType )
// 	if err != nil{
// 		return nil, err
// 	}
// 	if err := encoder.EncodeFileByFileType(code.FileType,f, qrImg);err != nil{
// 		return nil, err
// 	}
// 	return
// }
