package warapi

import (
	"encoding/json"

	fxu "github.com/j33pguy/fxutils"
)

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
	worldLayout := fxu.ReadJsonFile(file)

	o := new(WorldLayout)
	json.Unmarshal(*worldLayout, o)

	return o
}
