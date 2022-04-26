package fxapi

import (
	fxu "github.com/j33pguy/fxutils"
)

// DONE 4/25/22
type WorldMaps []string

func GetWorldMaps() *WorldMaps {
	param := "/worldconquest/maps"

	// ex generic func
	// funcSig[type](xxxxxx){}
	res := fxu.GetFXApi[WorldMaps](param)

	return res
}
