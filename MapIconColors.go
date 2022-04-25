package fxapi

import (
	"encoding/json"

	fxu "github.com/j33pguy/fxutils"
)

// DONE
type MapIconColors struct {
	FactionColors []struct {
		FactionName string   `json:"FactionName"`
		Rgbf        []string `json:"RGBF"`
		Rgb         []string `json:"RGB"`
		HexLinear   string   `json:"HexLinear"`
		HexSRGB     string   `json:"HexSRGB"`
	} `json:"FactionColors"`
}

func (s MapIconColors) GetMapIconColors() *MapIconColors {
	file := "MapIconColors.json"
	mapIconColors := fxu.ReadJsonFile(file)

	o := new(MapIconColors)
	json.Unmarshal(*mapIconColors, o)

	return o
}
