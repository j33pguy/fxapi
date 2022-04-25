package warapi

import (
	"encoding/json"
	"time"
)

// DONE
type WorldConquest struct {
	WarID                string    `json:"warId"`
	WarNumber            int       `json:"warNumber"`
	Winner               string    `json:"winner"`
	ConquestStartTime    time.Time `json:"conquestStartTime"`
	ConquestEndTime      time.Time `json:"conquestEndTime"`
	ResistanceStartTime  time.Time `json:"resistanceStartTime"`
	RequiredVictoryTowns int       `json:"requiredVictoryTowns"`
}

func (s WorldConquest) GetWorldConquest() *WorldConquest {
	param := "/worldconquest/war"
	baseurl := Baseurl(param)
	resp := Response(*baseurl)

	o := new(WorldConquest)
	json.Unmarshal(*resp, o)

	return o
}
