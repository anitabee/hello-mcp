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

func GetForecast(ctx context.Context, req *mcp.CallToolRequest, input ForecastInput) (*mcp.CallToolResult, ForecastOutput, error) {

	url := fmt.Sprintf("%s/points/%s,%s", NwsApiBase, input.Latitude, input.Longitude)

	pointsData, err := MakeNewRequest(url)
	if err != nil || pointsData == nil {
		e := fmt.Errorf("Something went wrong with making pointsData request: %v", err)
		return nil, ForecastOutput{}, e
	}

	var dataWp WeatherPoint
	err = json.Unmarshal([]byte(pointsData), &dataWp)
	if err != nil {
		e := fmt.Errorf("Error unmarshaling WeatherPoint JSON: %v", err)
		return nil, ForecastOutput{}, e
	}

	if dataWp.Properties.Forecast == "" {
		e := fmt.Errorf("No forecast URL found in points data")
		return nil, ForecastOutput{}, e
	}

	forecastData, err := MakeNewRequest(dataWp.Properties.Forecast)
	if err != nil || forecastData == nil {
		e := fmt.Errorf("Something went wrong with making forecastData request: %v", err)
		return nil, ForecastOutput{}, e
	}

	var dataWf WeatherForecast
	err = json.Unmarshal([]byte(forecastData), &dataWf)
	if err != nil {
		e := fmt.Errorf("Error unmarshaling forecast JSON: %v", err)
		return nil, ForecastOutput{}, e
	}
	if dataWf.Properties.Periods == nil || len(dataWf.Properties.Periods) == 0 {
		e := fmt.Errorf("No forecast periods found")
		return nil, ForecastOutput{}, e
	}

	formatForecasts := []string{}
	for _, period := range dataWf.Properties.Periods[:5] {
		formatted := fmt.Sprintf(`
		Name: %s
		Detailed Forecast: %s
		Temperature: %d°%s
		Wind: %s %s
		`, period.Name, period.DetailedForecast, period.Temperature, period.TemperatureUnit, period.WindSpeed, period.WindDirection)
		formatForecasts = append(formatForecasts, formatted)
	}
	formatedPeriods := strings.Join(formatForecasts, "\n---\n")
	return nil, ForecastOutput{Forecast: formatedPeriods}, nil

}
