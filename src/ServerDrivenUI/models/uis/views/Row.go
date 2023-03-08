package views

import (
	"ServerDrivenUI/models/uis/properties"
	"ServerDrivenUI/models/uis/types/viewgrouptype"
)

type row struct {
	Type     string        `json:"type" binding:"required"`
	Children []interface{} `json:"children"`
	Value    RowViewValue  `json:"rowViewValue,omitempty"`
}

func Row(rowViewValue RowViewValue, children []interface{}) *row {
	row := new(row)
	row.Type = viewgrouptype.ROW.String()
	row.Children = children
	row.Value = rowViewValue
	return row
}

type RowViewValue struct {
	Width           int              `json:"width,omitempty"`
	Height          int              `json:"height,omitempty"`
	BackgroundColor properties.Color `json:"backgroundColor,omitempty"`
	BorderColor     properties.Color `json:"borderColor,omitempty"`
	BorderWidth     int              `json:"borderWidth,omitempty"`
	PaddingL        int              `json:"paddingL,omitempty"`
	PaddingT        int              `json:"paddingT,omitempty"`
	PaddingR        int              `json:"paddingR,omitempty"`
	PaddingB        int              `json:"paddingB,omitempty"`
	Corner          string           `json:"cornerType,omitempty"`
	Radius          int              `json:"radius,omitempty"`
	ColumnAlignment string           `json:"columnAlignment,omitempty"`
}
