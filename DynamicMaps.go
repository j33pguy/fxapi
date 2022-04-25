package fxapi

import (
	"encoding/json"
)

// DONE
type DynamicMap struct {
	RegionID             int `json:"regionId"`
	ScorchedVictoryTowns int `json:"scorchedVictoryTowns"`
	MapItems             []struct {
		TeamID   string  `json:"teamId"`
		IconType int     `json:"iconType"`
		X        float64 `json:"x"`
		Y        float64 `json:"y"`
		Flags    int     `json:"flags"`
	} `json:"mapItems"`
	MapTextItems []interface{} `json:"mapTextItems"`
	LastUpdated  int64         `json:"lastUpdated"`
	Version      int           `json:"version"`
}

func (s DynamicMap) GetDynamicMap(mapName string) *DynamicMap {
	param := "/worldconquest/maps/" + mapName + "/dynamic/public"
	baseurl := Baseurl(param)
	resp := Response(*baseurl)

	o := new(DynamicMap)
	json.Unmarshal(*resp, o)

	return o
}
