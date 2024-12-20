package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/vladimirvivien/gexe"
)

// This example uses gexe HTTP support to retrieve
// Weather data from NOAA
func main() {
	exe := gexe.SetVar("lat", "40.7501").
		SetVar("long", "-111.482")

	// pull weather.go grid points data
	points := exe.Get("https://api.weather.gov/points/${lat},${long}").Body()
	defer points.Close()

	// decode grid points pointsData
	var pointsData map[string]any
	if err := json.NewDecoder(points).Decode(&pointsData); err != nil {
		fmt.Println("Unable to decode grid points data:", err)
		os.Exit(1)
	}

	// extract forecast url from points data above
	exe.SetVar("forecastUrl", pointsData["properties"].(map[string]any)["forecast"].(string))

	// retrieve forecast data from grid points above
	forecasts := exe.Get("$forecastUrl").Body()
	defer forecasts.Close()

	var forecastData map[string]any
	if err := json.NewDecoder(forecasts).Decode(&forecastData); err != nil {
		fmt.Println("Unable to decode forecast data:", err)
		os.Exit(1)
	}

	// loop through and display forecast periods
	forecastPeriods := forecastData["properties"].(map[string]any)["periods"].([]any)
	for _, period := range forecastPeriods {
		forecast := period.(map[string]any)
		fmt.Printf("%s: %s\n", forecast["name"], forecast["detailedForecast"])
	}
}
