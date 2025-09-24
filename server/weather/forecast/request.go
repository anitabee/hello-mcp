package forecast

import (
	"io"
	"net/http"
	"time"
)

const NWSAPIBase = "https://api.weather.gov"

func MakeNewRequest(url string) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "weather-app/1.0")
	req.Header.Set("Accept", "application/geo+json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return bodyBytes, nil
}
