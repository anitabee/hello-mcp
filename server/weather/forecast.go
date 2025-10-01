package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// Weather Forecast Structures
type ForecastPoint struct {
	Properties ForecastPointProperties `json:"properties"`
}

type ForecastPointProperties struct {
	Forecast string `json:"forecast"`
}

type Forecast struct {
	Properties ForecastProperties `json:"properties"`
}

type ForecastProperties struct {
	Periods []ForecastPeriod `json:"periods"`
}

type ForecastPeriod struct {
	Name             string `json:"name"`
	Temperature      int    `json:"temperature"`
	TemperatureUnit  string `json:"temperatureUnit"`
	WindSpeed        string `json:"windSpeed"`
	WindDirection    string `json:"windDirection"`
	DetailedForecast string `json:"detailedForecast"`
}

type GetForecastInput struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type GetForecastOutput struct {
	Forecast string `json:"forecast"`
}

func getForecastURL(input GetForecastInput) (string, error) {
	url := fmt.Sprintf("%s/points/%s,%s", NWSAPIBase, input.Latitude, input.Longitude)
	pointsData, err := makeNewRequest(url)
	if err != nil || pointsData == nil {
		return "", fmt.Errorf("something went wrong with making pointsData request: %v", err)
	}

	var dataWp ForecastPoint
	err = json.Unmarshal([]byte(pointsData), &dataWp)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling WeatherPoint JSON: %v", err)
	}

	if dataWp.Properties.Forecast == "" {
		return "", fmt.Errorf("no forecast URL found in points data")
	}
	return dataWp.Properties.Forecast, nil

}

func getForecast(ctx context.Context, req *mcp.CallToolRequest, input GetForecastInput) (*mcp.CallToolResult, GetForecastOutput, error) {

	ForecastURL, err := getForecastURL(input)
	if err != nil {
		return nil, GetForecastOutput{}, fmt.Errorf("error getting forecast URL: %v", err)
	}

	forecastData, err := makeNewRequest(ForecastURL)
	if err != nil || forecastData == nil {
		return nil, GetForecastOutput{}, fmt.Errorf("something went wrong with making forecastData request: %v", err)
	}

	var dataWf Forecast
	err = json.Unmarshal([]byte(forecastData), &dataWf)
	if err != nil {
		return nil, GetForecastOutput{}, fmt.Errorf("error unmarshaling forecast JSON: %v", err)
	}
	if len(dataWf.Properties.Periods) == 0 {
		return nil, GetForecastOutput{}, fmt.Errorf("no forecast periods found")
	}

	formatForecasts := []string{}
	for _, period := range dataWf.Properties.Periods[:5] {
		formatted := formatPerod(&period)
		formatForecasts = append(formatForecasts, formatted)
	}
	formatedPeriods := strings.Join(formatForecasts, "\n---\n")
	return nil, GetForecastOutput{Forecast: formatedPeriods}, nil

}

func formatPerod(period *ForecastPeriod) string {
	return fmt.Sprintf(`
		Name: %s
		Detailed Forecast: %s
		Temperature: %d°%s
		Wind: %s %s
		`,
		period.Name,
		period.DetailedForecast,
		period.Temperature,
		period.TemperatureUnit,
		period.WindSpeed,
		period.WindDirection,
	)
}
