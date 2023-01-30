package types

type View int

const (
	TEXT View = iota
	IMAGE
	EDIT_TEXT
)

func (view View) String() string {
	return [...]string{"TEXT", "IMAGE", "EDIT_TEXT"}[view]
}

func (view View) indexOf() int {
	return int(view)
}
