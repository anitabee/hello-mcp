package alerts

type WeatherAlertResponse struct {
	Features []Feature `json:"features"`
}

type Feature struct {
	Properties Properties `json:"properties"`
}

type Properties struct {
	AreaDesc    string `json:"areaDesc"`
	Severity    string `json:"severity"`
	Event       string `json:"event"`
	Description string `json:"description"`
	Instruction string `json:"instruction"`
}
