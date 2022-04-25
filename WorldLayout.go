package warapi

import (
	"encoding/json"

	"github.com/j33pguy/fxterminal/utils"
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
	file := "/home/j33p/Projects/J33PGUY/FoxholeTerminal/warapi/WorldLayout.json"
	worldLayout := utils.ReadJsonFile(file)

	o := new(WorldLayout)
	json.Unmarshal(*worldLayout, o)

	return o
}
