package fxapi

import (
	"time"

	fxu "github.com/j33pguy/fxutils"
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

func GetWorldConquest() *WorldConquest {
	param := "/worldconquest/war"
	res := fxu.GetFuncResourceApi[WorldConquest](param)

	return res
}
