package main

import (
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func registerMcpTools() {
	server := mcp.NewServer(&mcp.Implementation{Name: "weather", Version: "v0.0.1"}, nil)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_alerts",
		Description: "Get the current weather for a given location",
	}, getAlerts)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_forecast",
		Description: "Get the weather forecast for a given location",
	}, getForecast)

	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}

func main() {
	registerMcpTools()
}
