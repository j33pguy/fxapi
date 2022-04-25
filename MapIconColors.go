package warapi

import (
	"encoding/json"

	"github.com/j33pguy/fxterminal/utils"
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
	file := "/home/j33p/Projects/J33PGUY/FoxholeTerminal/warapi/MapIconColors.json"
	mapIconColors := utils.ReadJsonFile(file)

	o := new(MapIconColors)
	json.Unmarshal(*mapIconColors, o)

	return o
}
