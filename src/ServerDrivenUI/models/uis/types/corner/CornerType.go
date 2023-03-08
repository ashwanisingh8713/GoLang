package corner

type CornerType int

const (
	ROUNDED_CORNER CornerType = iota
	CUT_CORNER
)

func (cornerType CornerType) String() string {
	return [...]string{"ROUNDED_CORNER", "CUT_CORNER"}[cornerType]
}

func (cornerType CornerType) indexOf() int {
	return int(cornerType)
}
