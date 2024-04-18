package main

import (
	"qrcode_generator/app"
	"qrcode_generator/app/config"
)

func main() {
	app.Start(config.Config.Port)
}
