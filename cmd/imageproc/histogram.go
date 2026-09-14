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
