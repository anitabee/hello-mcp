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
