package warapi

import (
	"encoding/json"
)

// DONE
type WorldMaps []string

func (s WorldMaps) GetWorldMaps() *WorldMaps {
	param := "/worldconquest/maps"
	baseurl := Baseurl(param)
	resp := Response(*baseurl)

	o := new(WorldMaps)
	json.Unmarshal(*resp, o)

	return o
}
