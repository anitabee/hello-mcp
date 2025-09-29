package forecast

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type ForecastInput struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type ForecastOutput struct {
	Forecast string `json:"forecast"`
}

func getForecastURL(input ForecastInput) (string, error) {
	url := fmt.Sprintf("%s/points/%s,%s", NWSAPIBase, input.Latitude, input.Longitude)
	pointsData, err := MakeNewRequest(url)
	if err != nil || pointsData == nil {
		return "", fmt.Errorf("something went wrong with making pointsData request: %v", err)
	}

	var dataWp WeatherPoint
	err = json.Unmarshal([]byte(pointsData), &dataWp)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling WeatherPoint JSON: %v", err)
	}

	if dataWp.Properties.Forecast == "" {
		return "", fmt.Errorf("no forecast URL found in points data")
	}
	return dataWp.Properties.Forecast, nil

}

func GetForecast(ctx context.Context, req *mcp.CallToolRequest, input ForecastInput) (*mcp.CallToolResult, ForecastOutput, error) {

	ForecastURL, err := getForecastURL(input)
	if err != nil {
		return nil, ForecastOutput{}, fmt.Errorf("error getting forecast URL: %v", err)
	}

	forecastData, err := MakeNewRequest(ForecastURL)
	if err != nil || forecastData == nil {
		return nil, ForecastOutput{}, fmt.Errorf("something went wrong with making forecastData request: %v", err)
	}

	var dataWf WeatherForecast
	err = json.Unmarshal([]byte(forecastData), &dataWf)
	if err != nil {
		e := fmt.Errorf("error unmarshaling forecast JSON: %v", err)
		return nil, ForecastOutput{}, e
	}
	if len(dataWf.Properties.Periods) == 0 {
		e := fmt.Errorf("no forecast periods found")
		return nil, ForecastOutput{}, e
	}

	formatForecasts := []string{}
	for _, period := range dataWf.Properties.Periods[:5] {
		formatted := formatPerod(&period)
		formatForecasts = append(formatForecasts, formatted)
	}
	formatedPeriods := strings.Join(formatForecasts, "\n---\n")
	return nil, ForecastOutput{Forecast: formatedPeriods}, nil

}

func formatPerod(period *WeatherForecastPeriod) string {
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
