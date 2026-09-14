package imageproc

func InvertImage(img *LinearImage) {
	for i := range img.R {
		img.R[i] = 1.0 - img.R[i]
		img.G[i] = 1.0 - img.G[i]
		img.B[i] = 1.0 - img.B[i]
	}
}
