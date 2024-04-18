package handler

import (
	"bytes"
	"fmt"
	"image"
	"mime/multipart"
	"net/http"
	"qrcode_generator/app/utils/converter"
	"qrcode_generator/app/utils/encoder"
	"qrcode_generator/app/utils/file"
	"qrcode_generator/app/utils/generator"
	"strings"

	"github.com/gin-gonic/gin"
)

func GenerateQRCode(c *gin.Context) {
	var i struct {
		Content            string                `form:"content" json:"content" binding:"required"`
		Size               int                   `form:"size" json:"size" binding:""`
		IsSaveToFile       bool                  `form:"save" json:"save" binding:"boolean"`
		FileName           string                `form:"filename" json:"filename" binding:""`
		FileType           string                `form:"filetype" json:"filetype" binding:""`
		IsWatermarkEnabled bool                  `form:"watermark" json:"watermark" binding:"boolean"`
		WatermarkFile      *multipart.FileHeader `form:"watermarkfile" json:"watermarkfile" binding:""`
	}

	if err := c.ShouldBind(&i); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if i.Size < 45 {
		i.Size = 200
	}

	if i.IsSaveToFile {
		if strings.TrimSpace(i.FileType) == "" {
			i.FileType = "png"
		}
	}
	qrCodeData := generator.QRCodeData{Content: i.Content, Size: i.Size, HasWatermark: i.IsWatermarkEnabled, WatermarkImg: nil}

	if i.WatermarkFile != nil {
		watermarkFile, err := (i.WatermarkFile).Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		watermarkImg, _, err := image.Decode(watermarkFile)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		qrCodeData.WatermarkImg = watermarkImg
	}

	qrCode, err := qrCodeData.Generate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if i.IsSaveToFile {
		filePath := "resources/qrcode"
		f, err := file.SaveImage(qrCode, filePath, "qrcode", i.FileType)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.FileAttachment(f.Name(), fmt.Sprintf("%s.png", i.FileName))
		return
	}
	if !i.IsSaveToFile {
		buf := new(bytes.Buffer)
		if err := encoder.EncodeFileByFileType(i.FileType, buf, qrCode); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		qrCodeHTML, err := converter.ConvertImageToHTMLString(buf.Bytes())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(qrCodeHTML))
	}

}

func DownloadQRCodeSaved(c *gin.Context) {
	var i struct {
		FileName string `form:"filename" binding:""`
		FilePath string `form:"filepath" binding:"required"`
	}

	if err := c.ShouldBind(&i); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.FileAttachment(i.FilePath, fmt.Sprintf("%s.png", i.FileName))

}
