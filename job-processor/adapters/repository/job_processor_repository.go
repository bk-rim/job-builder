package repository

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/bk-rim/job-processor/domain"
)

type JobProcessorRepository struct {
}

func NewJobProcessorRepository() *JobProcessorRepository {
	return &JobProcessorRepository{}
}

func (r *JobProcessorRepository) UpadateJobStatus(job *domain.Job) error {

	var wg sync.WaitGroup
	wg.Add(1)

	errChan := make(chan error)

	var host string
	if os.Getenv("DOCKER_ENV") == "true" {
		host = "http://job-service:8080/jobs/" + strconv.Itoa(job.ID)
	} else {
		host = "http://localhost:8080/jobs/" + strconv.Itoa(job.ID)
	}

	go func() {
		defer wg.Done()

		jobFormated, err := json.Marshal(job)
		if err != nil {
			errChan <- err
			return
		}
		req, err := http.NewRequest("PUT", host, bytes.NewBuffer([]byte(jobFormated)))
		if err != nil {
			errChan <- err
			return
		}

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			errChan <- err
			return
		}
		defer resp.Body.Close()
	}()

	go func() {
		wg.Wait()
		close(errChan)
	}()

	return <-errChan
}
