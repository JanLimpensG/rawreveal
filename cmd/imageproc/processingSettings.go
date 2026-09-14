package imageproc

type ProcessingSettings struct {
	Exposure float32 `json:"exposure"`
	Invert   bool    `json:"invert"`
}

func SetDefaultProcessingSettings() ProcessingSettings {
	return ProcessingSettings{
		Exposure: 0.0,
		Invert:   false,
	}
}
