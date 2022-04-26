package fxapi

import (
	"encoding/json"

	fxu "github.com/j33pguy/fxutils"
)

// DONE 4/25/22
type MapIconColors struct {
	FactionColors []struct {
		FactionName string   `json:"FactionName"`
		Rgbf        []string `json:"RGBF"`
		Rgb         []string `json:"RGB"`
		HexLinear   string   `json:"HexLinear"`
		HexSRGB     string   `json:"HexSRGB"`
	} `json:"FactionColors"`
}

/// >>>> OLD FILE TO RIGHT

// these should all be local files in this dir
// original file im trying to parse from
func (s MapIconColors) GetMapIconColors() *MapIconColors {
	file := "MapIconColors.json"
	mapIconColors := fxu.GetLocalJson[MapIconColors](file)

	o := new(MapIconColors)
	json.Unmarshal(*mapIconColors, o)

	return o
}
