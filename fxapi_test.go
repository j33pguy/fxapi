package fxapi_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/j33pguy/fxutils"
	"github.com/stretchr/testify/assert"

	. "/fxapi" // Replace with the actual path to the fxapi package
)

func TestGetDynamicMap(t *testing.T) {
	// Define a sample dynamic map response
	dynamicMapJSON := `{
		"regionId": 1,
		"scorchedVictoryTowns": 2,
		"mapItems": [
			{
				"teamId": "A",
				"iconType": 3,
				"x": 1.2,
				"y": 2.3,
				"flags": 4
			}
		],
		"mapTextItems": [],
		"lastUpdated": 1234567890,
		"version": 1
	}`

	// Create a test server to mock API calls
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(dynamicMapJSON))
	}))
	defer ts.Close()

	// Replace the actual API URL with the test server URL
	fxutils.FXApiBaseURL = ts.URL

	// Call the GetDynamicMap function with a sample map name
	mapName := "sample_map"
	result := GetDynamicMap(mapName)

	// Unmarshal the expected response JSON into a DynamicMap struct
	var expected DynamicMap
	err := json.Unmarshal([]byte(dynamicMapJSON), &expected)
	assert.NoError(t, err)

	// Compare the result with the expected value
	assert.Equal(t, &expected, result)
}
