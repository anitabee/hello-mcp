package forecast

type WeatherPoint struct {
	Properties WeatherPointProperties `json:"properties"`
}

type WeatherPointProperties struct {
	Forecast string `json:"forecast"`
}

type WeatherForecast struct {
	Properties WeatherForecastProperties `json:"properties"`
}

type WeatherForecastProperties struct {
	Periods []WeatherForecastPeriod `json:"periods"`
}

type WeatherForecastPeriod struct {
	Name             string `json:"name"`
	Temperature      int    `json:"temperature"`
	TemperatureUnit  string `json:"temperatureUnit"`
	WindSpeed        string `json:"windSpeed"`
	WindDirection    string `json:"windDirection"`
	DetailedForecast string `json:"detailedForecast"`
}

type WeatherAlertResponse struct {
	Features []WeatherAlertResponseFeature `json:"features"`
}

type WeatherAlertResponseFeature struct {
	Properties WeatherAlertResponseProperties `json:"properties"`
}

type WeatherAlertResponseProperties struct {
	AreaDesc    string `json:"areaDesc"`
	Severity    string `json:"severity"`
	Event       string `json:"event"`
	Description string `json:"description"`
	Instruction string `json:"instruction"`
}
