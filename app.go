package main

import (
	"context"
	"rawreveal/cmd/imageproc"
)

// App struct
type App struct {
	ctx      context.Context
	document *imageproc.Document
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		document: &imageproc.Document{},
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) LoadImage(path string) (*imageproc.RenderResult, error) {
	err := a.document.Load(path)
	if err != nil {
		return nil, err
	}
	return a.document.RenderBase64()
}
