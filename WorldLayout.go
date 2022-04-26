package fxapi

import (
	"encoding/json"

	fxu "github.com/j33pguy/fxutils"
)

// DONE 4/25/22
type WorldLayout struct {
	WorldLayout struct {
		Columns []struct {
			ColumnNumber int `json:"ColumnNumber"`
			ColumnMaps   []struct {
				MapName   string `json:"mapName"`
				HumanName string `json:"humanName"`
			} `json:"ColumnMaps"`
		} `json:"Columns"`
	} `json:"WorldLayout"`
}

func (s WorldLayout) GetWorldLayout() *WorldLayout {
	file := "WorldLayout.json"
	worldLayout := fxu.GetLocalJson[WorldLayout](file)

	o := new(WorldLayout)
	json.Unmarshal(*worldLayout, o)

	return o
}
