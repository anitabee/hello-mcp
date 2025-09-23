package tool

import (
	"context"
	"log"

	"github.com/anitabee/hello-mcp/server/weather/forecast"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func RegisterMcpTools() {
	server := mcp.NewServer(&mcp.Implementation{Name: "weather", Version: "v0.0.1"}, nil)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_alerts",
		Description: "Get the current weather for a given location",
	}, forecast.GetAlerts)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_forecast",
		Description: "Get the weather forecast for a given location",
	}, forecast.GetForecast)

	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
