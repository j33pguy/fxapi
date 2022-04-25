package warapi

import (
	"encoding/json"

	fxu "github.com/j33pguy/fxutils"
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
	file := "MapIcons.json"
	mapDat := fxu.ReadJsonFile(file)

	o := new(MapIcons)
	json.Unmarshal(*mapDat, o)

	return o
}
