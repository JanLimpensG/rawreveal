package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"os"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

type ImageData struct {
	Base64Image string          `json:"base64Image"`
	Data        []HistogramData `json:"histogram"`
}

type HistogramData struct {
	Channel string `json:"channel"`
	Data    []int  `json:"data"`
}

func (a *App) ReadFile(path string) (ImageData, error) {
	file, err := os.Open(path)
	if err != nil {
		return ImageData{}, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return ImageData{}, err
	}

	rHist, gHist, bHist, grayHist, err := generateHistograms(img)
	if err != nil {
		return ImageData{}, err
	}

	// Encode inverted image to PNG in buffer
	var buf bytes.Buffer
	err = png.Encode(&buf, img)
	if err != nil {
		return ImageData{}, err
	}

	// Return base64 of inverted image
	return ImageData{
		Base64Image: base64.StdEncoding.EncodeToString(buf.Bytes()),
		Data: []HistogramData{
			{Channel: "Red", Data: rHist},
			{Channel: "Green", Data: gHist},
			{Channel: "Blue", Data: bHist},
			{Channel: "Grayscale", Data: grayHist},
		},
	}, nil
}

func (a *App) InvertImage(base64Src string) (string, error) {
	srcBytes, err := base64.StdEncoding.DecodeString(base64Src)
	if err != nil {
		return "", err
	}
	src, _, err := image.Decode(bytes.NewReader(srcBytes))
	if err != nil {
		return "", err
	}

	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := src.At(x, y).RGBA()

			// RGBA() returns 16-bit values (0-65535), convert to 8-bit (0-255)
			dst.SetRGBA(x, y, color.RGBA{
				R: uint8(255 - r>>8),
				G: uint8(255 - g>>8),
				B: uint8(255 - b>>8),
				A: uint8(a >> 8),
			})
		}
	}

	// Encode inverted image to PNG in buffer
	var buf bytes.Buffer
	err = png.Encode(&buf, dst)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

func generateHistograms(src image.Image) ([]int, []int, []int, []int, error) {
	rHist := make([]int, 256)
	gHist := make([]int, 256)
	bHist := make([]int, 256)
	grayHist := make([]int, 256)

	bounds := src.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := src.At(x, y).RGBA()
			// Convert to 8-bit values
			r8 := uint8(r >> 8)
			g8 := uint8(g >> 8)
			b8 := uint8(b >> 8)
			gray := uint8((r8 + g8 + b8) / 3)

			rHist[r8]++
			gHist[g8]++
			bHist[b8]++
			grayHist[gray]++
		}
	}

	return rHist, gHist, bHist, grayHist, nil
}
