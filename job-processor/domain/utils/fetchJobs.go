package domain_utils

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/bk-rim/job-processor/domain"
)

func FetchJobs() ([]domain.Job, error) {
	var host string
	if os.Getenv("DOCKER_ENV") == "true" {
		host = "http://job-service:8080/jobs"
	} else {
		host = "http://localhost:8080/jobs"
	}
	resp, err := http.Get(host)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var jobs []domain.Job
	err = json.NewDecoder(resp.Body).Decode(&jobs)
	if err != nil {
		return nil, err
	}

	return jobs, nil
}
