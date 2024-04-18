package router

import (
	"qrcode_generator/app/handler"

	"github.com/gin-gonic/gin"
)

func SetUpRouters(g *gin.Engine) *gin.Engine {

	qr := g.Group("")
	qr.POST("/generate", handler.GenerateQRCode)
	qr.GET("/download", handler.DownloadQRCodeSaved)
	return g
}
