package imgscale

type ImgScale int

const (
	Fit ImgScale = iota
	Inside
	FillBounds
	FillWidth
	FillHeight
	Crop
)

func (imgScale ImgScale) String() string {
	return [...]string{"Fit", "Inside", "FillBounds", "FillWidth", "FillHeight", "Crop"}[imgScale]
}

func (imgScale ImgScale) indexOf() int {
	return int(imgScale)
}
