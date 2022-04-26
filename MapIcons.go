package fxapi

import (
	"encoding/json"

	fxu "github.com/j33pguy/fxutils"
)

// DONE 4/25/22
type MapIcons struct {
	MapIcons []struct {
		IconName   string `json:"IconName"`
		IconNumber int    `json:"IconNumber"`
		IconNote   string `json:"IconNote"`
	} `json:"MapIcons"`
}

func (s MapIcons) GetMapIcons() *MapIcons {
	file := "MapIcons.json"
	mapDat := fxu.GetLocalJson[MapIcons](file)

	o := new(MapIcons)
	json.Unmarshal(*mapDat, o)

	return o
}
