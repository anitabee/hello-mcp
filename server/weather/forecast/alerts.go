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

	// log.Println("hello from GetAlerts where will my messages go?")

	url := fmt.Sprintf("%s/alerts/active/area/%s", NWSAPIBase, input.State)

	bodyBytes, err := MakeNewRequest(url)
	if err != nil || bodyBytes == nil {
		return nil, Output{}, fmt.Errorf("something went wrong with making new request: %v", err)
	}

	var data WeatherAlertResponse
	err = json.Unmarshal([]byte(bodyBytes), &data)
	if err != nil {
		return nil, Output{}, fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	if len(data.Features) == 0 {
		return nil, Output{}, fmt.Errorf("no alerts features found: %v", err)
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
