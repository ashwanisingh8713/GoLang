package alignment

type ColumnAlignment int

const (
	CenterHorizontally ColumnAlignment = iota
	Start
	End
)

func (columnAlignment ColumnAlignment) String() string {
	return [...]string{"CenterHorizontally", "Start", "End"}[columnAlignment]
}

func (columnAlignment ColumnAlignment) indexOf() int {
	return int(columnAlignment)
}
