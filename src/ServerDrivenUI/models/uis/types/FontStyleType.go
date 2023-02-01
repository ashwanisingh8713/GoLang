package types

type FontStyleType int

const (
	MontBold FontStyleType = iota
	MontSemiBold
	MontMedium
)

func (fontStyleType FontStyleType) String() string {
	return [...]string{"mont_bold", "mont_semibold", "mont_medium"}[fontStyleType]
}

func (fontStyleType FontStyleType) indexOf() int {
	return int(fontStyleType)
}
