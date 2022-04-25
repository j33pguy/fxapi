package warapi

import (
	"encoding/json"
)

// DONE
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

func (s StaticMap) GetStaticMap(mapName string) *StaticMap {
	param := "/worldconquest/maps/" + mapName + "/static"
	baseurl := Baseurl(param)
	resp := Response(*baseurl)

	o := new(StaticMap)
	json.Unmarshal(*resp, o)

	return o
}
