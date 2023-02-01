package views

import (
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/properties"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/types/viewgrouptype"
)

type horizontalGrid struct {
	Type       string              `json:"type" binding:"required"`
	Modifier   properties.Modifier `json:"modifier"`
	Children   []interface{}       `json:"children"`
	Value      properties.Value    `json:"value,omitempty"`
	GridColumn int                 `json:"gridColumn,omitempty"`
	GridHeight int                 `json:"gridHeight,omitempty"`
}

func HorizontalGrid(rowCount int, gridHeight int, modifier properties.Modifier, children []interface{}) *horizontalGrid {
	horizontalGrid := new(horizontalGrid)
	horizontalGrid.Type = viewgrouptype.HORIZONTAL_GRID.String()
	horizontalGrid.Modifier = modifier
	horizontalGrid.Children = children
	horizontalGrid.Value = properties.Value{Width: -1}
	horizontalGrid.GridColumn = rowCount
	horizontalGrid.GridHeight = gridHeight
	return horizontalGrid
}
