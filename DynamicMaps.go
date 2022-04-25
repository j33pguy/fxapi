package fxapi

import fxu "github.com/j33pguy/fxutils"

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

func (s StaticMap) GetStaticMap(mapName string) *StaticMap {
	param := "/worldconquest/maps/" + mapName + "/static"
	res := fxu.GetFuncResourceApi[StaticMap](param)

	return res
}
