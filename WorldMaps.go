package fxapi

import (
	fxu "github.com/j33pguy/fxutils"
)

// DONE
type WorldMaps []string

func GetWorldMaps() *WorldMaps {
	param := "/worldconquest/maps"
	res := fxu.GetFuncResourceApi[WorldMaps](param)

	return res
}
