package views

import (
	"ServerDrivenUI/models/uis/types/viewgrouptype"
)

type scaffold struct {
	Type     string        `json:"type" binding:"required"`
	Children []interface{} `json:"children""`
	APP_BAR  []interface{} `json:"top_bar"`
}

func Scaffold(children []interface{}, appBar []interface{}) *scaffold {
	scaffold := new(scaffold)
	scaffold.Type = viewgrouptype.SCAFFOLD.String()
	scaffold.Children = children
	scaffold.APP_BAR = appBar
	return scaffold
}
