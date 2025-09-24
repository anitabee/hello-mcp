# Hello MCP weather server

Implementation of the [Build an MCP server tutorial](https://modelcontextprotocol.io/docs/develop/build-server) in Go.  

It covers tools to fetch weather forecasts and alerts for specific locations using the [Weather Service API](https://www.weather.gov/documentation/services-web-api). It is designed to integrate with the Model Context Protocol (MCP) framework and is intended purely for my learning purposes.


## Features

- **Get Weather Forecast**: Fetches a detailed weather forecast for a given latitude and longitude.
- **Get Weather Alerts**: Retrieves active weather alerts for a given state.
- **MCP Integration**: Implements tools that can be registered and used with the MCP framework.

## Prerequisites

- Go 1.25.1 or later
- MCP Go SDK (`github.com/modelcontextprotocol/go-sdk`)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/anitabee/hello-mcp.git
   cd hello-mcp/server/weather
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Build the server:
   ```bash
   go build -o weather-server
   ```

   This will create an executable named `weather-server` in the current directory.

## Usage

### Running the Server

```bash
./weather-server
```

This registers the `get_forecast` and `get_alerts` tools and starts the server.  
In most cases, you would integrate with a client such as Claude, which handles starting the server for you (see [Configuration](#configuration))

### Example Tool: Get Weather Forecast

The `get_forecast` tool fetches the weather forecast for a given latitude and longitude.

Example input:
```json
{
  "latitude": "40.7128",
  "longitude": "-74.0060"
}
```

### Example Tool: Get Weather Alerts

The `get_alerts` tool retrieves active weather alerts for a given state.

Example input:
```json
{
  "state": "NY"
}
```

Example output:
```json
{
  "alerts": [
    {
      "title": "Severe Thunderstorm Warning",
      "description": "A severe thunderstorm has been detected in your area."
    },
    {
      "title": "Flood Watch",
      "description": "Heavy rainfall is expected, leading to potential flooding."
    }
  ]
}
```

## Configuration

The MCP server configuration is defined in `claude_desktop_config.json`. For example:

```json
{
  "mcpServers": {
    "weather": {
      "command": "/Users/<user-workspace>/hello-mcp/server/weather/weather-server"
    }
  }
}
```
