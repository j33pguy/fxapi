package fxapi

import (
	fxu "github.com/j33pguy/fxutils"
)

// DONE 4/25/22
type StaticMap struct {
	RegionID             int           `json:"regionId"`
	ScorchedVictoryTowns int           `json:"scorchedVictoryTowns"`
	MapItems             []interface{} `json:"mapItems"`
	MapTextItems         []struct {
		Text          string  `json:"text"`
		X             float64 `json:"x"`
		Y             float64 `json:"y"`
		MapMarkerType string  `json:"mapMarkerType"`
	} `json:"mapTextItems"`
	LastUpdated int64 `json:"lastUpdated"`
	Version     int   `json:"version"`
}

func GetStaticMap(mapName string) *StaticMap {
	param := "/worldconquest/maps/" + mapName + "/static"
	res := fxu.GetFXApi[StaticMap](param)

	return res
}
