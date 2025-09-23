package request

import (
	"io"
	"net/http"
	"time"
)

const (
	NwsApiBase = "https://api.weather.gov"
	UserAgent  = "weather-app/1.0"
)

func MakeNewRequest(url string) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", NwsApiBase)
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
