package main

import (
	"context"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type AppService struct{}

func (g *AppService) ChangeWindowOptions(options string, ctx context.Context) {
	app := application.Get()
	window := app.Window.Current()
	if options == "mini" {
		window.SetSize(200, 500)
		window.SetBackgroundColour(application.RGBA{Red: 0, Green: 0, Blue: 0, Alpha: 0})
	} else {
		window.SetSize(1200, 768)
		window.SetBackgroundColour(application.RGBA{Red: 255, Green: 255, Blue: 255, Alpha: 1})
	}
}
