package warapi

import (
	"encoding/json"
)

// DONE
type WarReport struct {
	TotalEnlistments   int `json:"totalEnlistments"`
	ColonialCasualties int `json:colonialCasualties"`
	WardenCasualties   int `json:wardenCasualties"`
	DayOfWar           int `json:"dayOfWar"`
}

func (s WarReport) GetWarReport(mapName string) *WarReport {
	param := "/worldconquest/warReport/" + mapName
	baseurl := Baseurl(param)
	resp := Response(*baseurl)

	o := new(WarReport)
	json.Unmarshal(*resp, o)

	return o
}
