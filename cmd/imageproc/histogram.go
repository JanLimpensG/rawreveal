package imageproc

type Histogram struct {
	R    [256]uint64 `json:"r"`
	G    [256]uint64 `json:"g"`
	B    [256]uint64 `json:"b"`
	Gray [256]uint64 `json:"gray"`
}

func NewHistogram() *Histogram {
	return &Histogram{}
}

func (h *Histogram) AddPixel(r, g, b uint8) {
	h.R[r]++
	h.G[g]++
	h.B[b]++

	gray := uint8(0.2126*float64(r) +
		0.7152*float64(g) +
		0.0722*float64(b),
	)

	h.Gray[gray]++
}

func CalculateHistogram(img *LinearImage) Histogram {
	hist := NewHistogram()

	for i := range img.R {
		r := Clamp01(LinearToSRGB(img.R[i]))
		g := Clamp01(LinearToSRGB(img.G[i]))
		b := Clamp01(LinearToSRGB(img.B[i]))

		ri := uint8(r*255 + 0.5)
		gi := uint8(g*255 + 0.5)
		bi := uint8(b*255 + 0.5)
		hist.AddPixel(ri, gi, bi)
	}
	return *hist
}
