package app

import (
	"qrcode_generator/app/router"

	"github.com/gin-gonic/gin"
)

type (
	Host struct {
		Gin *gin.Engine
	}
)

func Start(port string) {
	g := gin.New()

	g.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	g.ForwardedByClientIP = true

	g.SetTrustedProxies([]string{"*"})

	router.SetUpRouters(g)

	g.Run(":" + port)
}
