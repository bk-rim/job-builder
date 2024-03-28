package domain_utils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bk-rim/job-processor/domain"
)

func TestPostOnWebhookSlack(t *testing.T) {

	stubServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var slackMessage SlackMessage
		err := json.NewDecoder(r.Body).Decode(&slackMessage)
		if err != nil {
			t.Errorf("Failed to decode request payload: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))

	defer stubServer.Close()

	job := domain.Job{
		Type:         "weather",
		WebhookSlack: stubServer.URL,
	}

	data := `{"current":{"Temperature2m":25,"RelativeHumidity2m":60}}`

	status, err := PostOnWebhookSlack(data, job)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if status != "200" {
		t.Errorf("Expected status code 200, but got %s", status)
	}
}
