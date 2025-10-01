package main

import (
	"fmt"
	"testing"
)

func TestGetForecast_Success(t *testing.T) {
	input := ForecastInput{
		Latitude:  "40.7128",
		Longitude: "-74.0060",
	}
	_, out, _ := getForecast(nil, nil, input)

	fmt.Printf("Output: %+v\n", out)
}
