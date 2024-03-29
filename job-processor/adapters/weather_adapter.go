package adapter

import (
	"io/ioutil"
	"net/http"
	"sync"
)

type WeatherAPIAdapter struct {
}

func NewWeatherAPIAdapter() *WeatherAPIAdapter {
	return &WeatherAPIAdapter{}
}

func (a *WeatherAPIAdapter) FetchWeatherData() (string, error) {
	var wg sync.WaitGroup
	wg.Add(1)

	weatherData := make(chan string)
	errChan := make(chan error)

	go func() {
		defer wg.Done()

		resp, err := http.Get("https://api.open-meteo.com/v1/forecast?latitude=52.52&longitude=13.41&current=temperature_2m,relative_humidity_2m")
		if err != nil {
			errChan <- err
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			errChan <- err
			return
		}

		weatherData <- string(body)
	}()

	go func() {
		wg.Wait()
		close(weatherData)
		close(errChan)
	}()

	select {
	case status := <-weatherData:
		return status, nil
	case err := <-errChan:
		return "", err
	}
}
