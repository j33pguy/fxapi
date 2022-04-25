package warapi

import (
	"encoding/json"

	"github.com/j33pguy/fxterminal/utils"
)

// DONE
type MapIcons struct {
	MapIcons []struct {
		IconName   string `json:"IconName"`
		IconNumber int    `json:"IconNumber"`
		IconNote   string `json:"IconNote"`
	} `json:"MapIcons"`
}

func (s MapIcons) GetMapIcons() *MapIcons {
	file := "/home/j33p/Projects/J33PGUY/FoxholeTerminal/warapi/MapIcons.json"
	mapDat := utils.ReadJsonFile(file)

	o := new(MapIcons)
	json.Unmarshal(*mapDat, o)

	return o
}
