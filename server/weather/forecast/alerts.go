package forecast

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type Input struct {
	State string `json:"state"`
}

type Output struct {
	Alerts string `json:"alerts"`
}

func GetAlerts(ctx context.Context, req *mcp.CallToolRequest, input Input) (*mcp.CallToolResult, Output, error) {
	url := fmt.Sprintf("%s/alerts/active/area/%s", NwsApiBase, input.State)

	bodyBytes, err := MakeNewRequest(url)
	if err != nil || bodyBytes == nil {
		e := fmt.Errorf("Something went wrong with making new request: %v", err)
		return nil, Output{}, e
	}

	var data WeatherAlertResponse
	err = json.Unmarshal([]byte(bodyBytes), &data)
	if err != nil {
		e := fmt.Errorf("Error unmarshaling JSON: %v", err)
		return nil, Output{}, e
	}

	if data.Features == nil || len(data.Features) == 0 {
		e := fmt.Errorf("No alerts features found: %v", err)
		return nil, Output{}, e
	}

	formatAlerts := []string{}
	for _, alert := range data.Features {
		formatted := formatAlert(&alert.Properties)
		formatAlerts = append(formatAlerts, formatted)
	}

	alerts := strings.Join(formatAlerts, "\n---\n")

	return nil, Output{
		Alerts: alerts,
	}, nil
}

func formatAlert(ap *WeatherAlertResponseProperties) string {
	return fmt.Sprintf(`
		Event: %s
		Area: %s
		Severity: %s
		Description: %s
		Instructions: %s`,
		ap.Event,
		ap.AreaDesc,
		ap.Severity,
		ap.Description,
		ap.Instruction,
	)
}
