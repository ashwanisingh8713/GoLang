package views

import (
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/properties"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/types/viewgrouptype"
)

type verticalGrid struct {
	Type     string            `json:"type" binding:"required"`
	Children []interface{}     `json:"children"`
	Value    VerticalGridValue `json:"verticalGridValue,omitempty"`
}

func VerticalGrid(verticalGridValue VerticalGridValue, children []interface{}) *verticalGrid {
	verticalGrid := new(verticalGrid)
	verticalGrid.Type = viewgrouptype.VERTICAL_GRID.String()
	verticalGrid.Children = children
	verticalGrid.Value = verticalGridValue
	return verticalGrid
}

type VerticalGridValue struct {
	Width                 int              `json:"width,omitempty"`
	Height                int              `json:"height,omitempty"`
	PaddingL              int              `json:"paddingL,omitempty"`
	PaddingT              int              `json:"paddingT,omitempty"`
	PaddingR              int              `json:"paddingR,omitempty"`
	PaddingB              int              `json:"paddingB,omitempty"`
	BackgroundColor       properties.Color `json:"backgroundColor,omitempty"`
	Corner                string           `json:"cornerType,omitempty"`
	IsEnable              bool             `json:"isEnable,omitempty"`
	Radius                int              `json:"radius,omitempty"`
	GridColumn            int              `json:"gridColumn" binding:"required"`
	GridHeight            int              `json:"gridHeight" binding:"required"`
	VerticalArrangement   int              `json:"verticalArrangement,omitempty"`
	HorizontalArrangement int              `json:"horizontalArrangement,omitempty"`
}
