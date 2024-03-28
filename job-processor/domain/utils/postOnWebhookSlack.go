package domain_utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/bk-rim/job-processor/domain"
)

type CurrentData struct {
	Time               string  `json:"time"`
	Interval           int     `json:"interval"`
	Temperature2m      float64 `json:"temperature_2m"`
	RelativeHumidity2m float64 `json:"relative_humidity_2m"`
}

type SlackMessage struct {
	Text string `json:"text"`
}

func PostOnWebhookSlack(data string, job domain.Job) (string, error) {

	var wg sync.WaitGroup
	wg.Add(1)

	errChan := make(chan error)
	statusChan := make(chan string)

	go func() {
		defer wg.Done()
		var slackText string

		switch job.Type {
		case "weather":
			var currentData struct {
				Current CurrentData `json:"current"`
			}
			err := json.Unmarshal([]byte(data), &currentData)
			if err != nil {
				errChan <- err
				return
			}
			slackText = "la température actuelle est de " + fmt.Sprintf("%v", currentData.Current.Temperature2m) + "°C et l'humidité relative est de " + fmt.Sprintf("%v", currentData.Current.RelativeHumidity2m) + "%"
		case "bridge_status":
			slackText = data
		}
		slackMessage := SlackMessage{Text: slackText}
		slackMessageJSON, err := json.Marshal(slackMessage)
		if err != nil {
			errChan <- err
			return
		}

		resp, err := http.Post(job.WebhookSlack, "application/json", bytes.NewBuffer(slackMessageJSON))
		if err != nil {
			errChan <- err
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			errChan <- fmt.Errorf("unexpected status code when posting on webhookSlack: %d", resp.StatusCode)
			statusChan <- strconv.Itoa(resp.StatusCode)
			return
		}
		statusChan <- strconv.Itoa(resp.StatusCode)
	}()

	go func() {
		wg.Wait()
		close(errChan)
		close(statusChan)
	}()

	select {
	case status := <-statusChan:
		return status, nil
	case err := <-errChan:
		return "", err
	}

}
