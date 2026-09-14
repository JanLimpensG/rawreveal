package imageproc

import (
	"errors"
	"image"
	"os"
)

type Document struct {
	OriginalImage *LinearImage       `json:"originalImage"`
	Settings      ProcessingSettings `json:"settings"`
}

func (d *Document) Load(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}
	// Process the decoded image as needed
	d.OriginalImage = ConvertToLinearImage(img)
	d.Settings = SetDefaultProcessingSettings()

	return nil
}

func (d *Document) Render() (image.Image, error) {
	if d.OriginalImage == nil {
		return nil, errors.New("no original image loaded")
	}

	working := d.OriginalImage.Clone()

	ApplySettings(working, d.Settings)

	return ConvertToSRGBImage(working), nil
}

func (d *Document) RenderBase64() (string, error) {
	img, err := d.Render()
	if err != nil {
		return "", err
	}

	return EncodeJPEGBase64(img, 90)
}
