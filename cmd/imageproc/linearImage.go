package imageproc

type LinearImage struct {
	Width  int
	Height int
	R      []float32
	G      []float32
	B      []float32
}

func (i *LinearImage) Index(x, y int) int {
	return y*i.Width + x
}

func (i *LinearImage) Clone() *LinearImage {
	clone := &LinearImage{
		Width:  i.Width,
		Height: i.Height,
		R:      make([]float32, len(i.R)),
		G:      make([]float32, len(i.G)),
		B:      make([]float32, len(i.B)),
	}
	copy(clone.R, i.R)
	copy(clone.G, i.G)
	copy(clone.B, i.B)
	return clone
}
