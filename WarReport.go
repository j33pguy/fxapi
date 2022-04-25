package fxapi

import (
	fxu "github.com/j33pguy/fxutils"
)

// DONE
type WarReport struct {
	TotalEnlistments   int `json:"totalEnlistments"`
	ColonialCasualties int `json:colonialCasualties"`
	WardenCasualties   int `json:wardenCasualties"`
	DayOfWar           int `json:"dayOfWar"`
}

func GetWarReport(mapName string) *WarReport {
	param := "/worldconquest/warReport/" + mapName
	res := fxu.GetFuncResourceApi[WarReport](param)

	return res
}
