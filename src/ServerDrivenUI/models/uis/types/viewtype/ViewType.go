package viewtype

type View int

const (
	TEXT View = iota
	IMAGE
	EDIT_TEXT
	WEBVIEW
)

func (view View) String() string {
	return [...]string{"TEXT", "IMAGE", "EDIT_TEXT", "WEBVIEW"}[view]
}

func (view View) indexOf() int {
	return int(view)
}
