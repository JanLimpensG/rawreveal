package imageproc

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/color"
	"image/jpeg"
	"math"
)

func Clamp01(v float32) float32 {
	if v < 0 {
		return 0
	}
	if v > 1 {
		return 1
	}
	return v
}

func SRGBToLinear(v float32) float32 {
	if v <= 0.04045 {
		return v / 12.92
	}
	return float32(math.Pow(float64((v+0.055)/1.055), 2.4))
}

func LinearToSRGB(v float32) float32 {
	if v <= 0.0031308 {
		return v * 12.92
	}
	return 1.055*float32(math.Pow(float64(v), 1/2.4)) - 0.055
}

func ConvertToLinearImage(img image.Image) *LinearImage {
	bounds := img.Bounds()
	imageWidth, imageHeight := bounds.Dx(), bounds.Dy()
	size := imageWidth * imageHeight
	linearImage := &LinearImage{
		Width:  imageWidth,
		Height: imageHeight,
		R:      make([]float32, size),
		G:      make([]float32, size),
		B:      make([]float32, size),
	}

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			i := linearImage.Index(x, y)
			linearImage.R[i] = SRGBToLinear(float32(r) / 65535.0)
			linearImage.G[i] = SRGBToLinear(float32(g) / 65535.0)
			linearImage.B[i] = SRGBToLinear(float32(b) / 65535.0)
		}
	}
	return linearImage
}

func ConvertToSRGBImage(img *LinearImage) image.Image {
	bounds := image.Rect(0, 0, img.Width, img.Height)
	srgbImage := image.NewRGBA(bounds)

	for y := 0; y < img.Height; y++ {
		for x := 0; x < img.Width; x++ {
			i := img.Index(x, y)
			r := Clamp01(LinearToSRGB(img.R[i]))
			g := Clamp01(LinearToSRGB(img.G[i]))
			b := Clamp01(LinearToSRGB(img.B[i]))

			srgbImage.SetRGBA(x, y, color.RGBA{
				R: uint8(r*255 + 0.5), // Adding 0.5 for rounding to the nearest integer
				G: uint8(g*255 + 0.5),
				B: uint8(b*255 + 0.5),
				A: 255,
			})
		}
	}
	return srgbImage
}

func EncodeJPEGBase64(img image.Image, quality int) (string, error) {
	var buf bytes.Buffer

	err := jpeg.Encode(&buf, img, &jpeg.Options{Quality: quality})
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

func ApplySettings(img *LinearImage, settings ProcessingSettings) {
	return // Placeholder for future implementation of processing settings
}
