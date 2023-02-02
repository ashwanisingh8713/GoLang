package views

import (
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/properties"
	"ServerDrivenUI/src/ServerDrivenUI/models/uis/types/viewgrouptype"
)

type verticalGrid struct {
	Type       string              `json:"type" binding:"required"`
	Modifier   properties.Modifier `json:"modifier"`
	Children   []interface{}       `json:"children"`
	Value      properties.Value    `json:"value,omitempty"`
	GridColumn int                 `json:"gridColumn,omitempty"`
	GridHeight int                 `json:"gridHeight,omitempty"`
}

func VerticalGrid(rowCount int, gridHeight int, modifier properties.Modifier, children []interface{}) *verticalGrid {
	verticalGrid := new(verticalGrid)
	verticalGrid.Type = viewgrouptype.VERTICAL_GRID.String()
	verticalGrid.Modifier = modifier
	verticalGrid.Children = children
	verticalGrid.Value = properties.Value{Width: -1}
	verticalGrid.GridColumn = rowCount
	verticalGrid.GridHeight = gridHeight
	return verticalGrid
}
