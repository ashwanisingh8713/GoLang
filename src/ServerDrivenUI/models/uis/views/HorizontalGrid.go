package views

import (
	"ServerDrivenUI/models/uis/properties"
	"ServerDrivenUI/models/uis/types/viewgrouptype"
)

type horizontalGrid struct {
	Type     string              `json:"type" binding:"required"`
	Children []interface{}       `json:"children"`
	Value    HorizontalGridValue `json:"horizontalGridValue,omitempty"`
}

func HorizontalGrid(horizontalGridValue HorizontalGridValue, children []interface{}) *horizontalGrid {
	horizontalGrid := new(horizontalGrid)
	horizontalGrid.Type = viewgrouptype.HORIZONTAL_GRID.String()
	horizontalGrid.Children = children
	horizontalGrid.Value = horizontalGridValue
	return horizontalGrid
}

type HorizontalGridValue struct {
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
