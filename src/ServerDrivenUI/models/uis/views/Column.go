package views

import (
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/properties"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/types/viewgrouptype"
)

type column struct {
	Type     string          `json:"type" binding:"required"`
	Children []interface{}   `json:"children"`
	Value    ColumnViewValue `json:"columnViewValue,omitempty"`
}

func Column(columnViewValue ColumnViewValue, children []interface{}) *column {
	column := new(column)
	column.Type = viewgrouptype.COLUMN.String()
	column.Children = children
	column.Value = columnViewValue
	return column
}

type ColumnViewValue struct {
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
