package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type AlertInput struct {
	State string `json:"state"`
}

type AlertOutput struct {
	Alerts string `json:"alerts"`
}

type AlertResponse struct {
	Features []AlertResponseFeature `json:"features"`
}

type AlertResponseFeature struct {
	Properties AlertResponseProperties `json:"properties"`
}

type AlertResponseProperties struct {
	AreaDesc    string `json:"areaDesc"`
	Severity    string `json:"severity"`
	Event       string `json:"event"`
	Description string `json:"description"`
	Instruction string `json:"instruction"`
}

func getAlerts(ctx context.Context, req *mcp.CallToolRequest, input AlertInput) (*mcp.CallToolResult, AlertOutput, error) {
	log.Println("hello from GetAlerts where will my messages go?")
	url := fmt.Sprintf("%s/alerts/active/area/%s", NWSAPIBase, input.State)

	bodyBytes, err := makeNewRequest(url)
	if err != nil || bodyBytes == nil {
		return nil, AlertOutput{}, fmt.Errorf("something went wrong with making new request: %v", err)
	}

	var data AlertResponse
	err = json.Unmarshal([]byte(bodyBytes), &data)
	if err != nil {
		return nil, AlertOutput{}, fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	if len(data.Features) == 0 {
		return nil, AlertOutput{}, fmt.Errorf("no alerts features found: %v", err)
	}

	formatAlerts := []string{}
	for _, alert := range data.Features {
		formatted := formatAlert(&alert.Properties)
		formatAlerts = append(formatAlerts, formatted)
	}

	alerts := strings.Join(formatAlerts, "\n---\n")

	return nil, AlertOutput{
		Alerts: alerts,
	}, nil
}

func formatAlert(ap *AlertResponseProperties) string {
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
