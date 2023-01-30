package types

type ViewGroup int

const (
	SCAFFOLD ViewGroup = iota
	APP_BAR
	VERTICAL_LIST
	HORIZONTAL_LIST
	ROW
	COLUMN
	BOX
	CARD
	GRID
	PAGER
)

func (viewGroup ViewGroup) String() string {
	return [...]string{"SCAFFOLD", "APP_BAR", "VERTICAL_LIST", "HORIZONTAL_LIST", "ROW", "COLUMN", "BOX", "CARD",
		"GRID", "PAGER"}[viewGroup]
}

func (viewGroup ViewGroup) indexOf() int {
	return int(viewGroup)
}
