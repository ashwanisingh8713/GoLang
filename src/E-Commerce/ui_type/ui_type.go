package ui_type

// ViewGroupType /**
type ViewGroupType int

const (
	HorizontalListView ViewGroupType = iota
	ViewPager
	Box
	HorizontalGrid
	VerticalGrid
)

func (viewGroupType ViewGroupType) String() string {
	return [...]string{"HorizontalListView", "ViewPager", "Box", "HorizontalGrid", "VerticalGrid"}[viewGroupType]
}

// TemplateViewType /**
type TemplateViewType int

const (
	Vt01 TemplateViewType = iota
	Vt02
	Vt03
	Vt04
	Vt05
	Vt06
	Vt07
	Vt08
	Vt09
	Vt10
	Vt11
	Vt12
	Vt13
)

func (viewType TemplateViewType) String() string {
	return [...]string{"Template_01", "Template_02", "Template_03", "Template_04", "Template_05", "Template_06",
		"Template_07", "Template_08", "Template_09", "Template_10", "Template_11", "Template_12", "Template_13"}[viewType]
}

// HeaderType /**
type HeaderType int

const (
	Header_01 HeaderType = iota
)

func (headerType HeaderType) String() string {
	return [...]string{"Header_01"}[headerType]
}
