package fxapi

import (
	fxu "github.com/j33pguy/fxutils"
)

// DONE 4/25/22
type WarReport struct {
	TotalEnlistments   int `json:"totalEnlistments"`
	ColonialCasualties int `json:colonialCasualties"`
	WardenCasualties   int `json:wardenCasualties"`
	DayOfWar           int `json:"dayOfWar"`
}

func GetWarReport(mapName string) *WarReport {
	param := "/worldconquest/warReport/" + mapName
	res := fxu.GetFXApi[WarReport](param)

	return res
}
