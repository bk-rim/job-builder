package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/bk-rim/job-service/domain"
)

type JobExecuterRepository struct {
}

func NewJobExecuterRepository() *JobExecuterRepository {
	return &JobExecuterRepository{}
}

func (r *JobExecuterRepository) LaunchJobExecution(job domain.Job) error {
	var wg sync.WaitGroup
	wg.Add(1)

	errChan := make(chan error)
	jobFormated, err := json.Marshal(job)
	if err != nil {
		return err
	}
	var host string
	if os.Getenv("DOCKER_ENV") == "true" {
		host = "http://job-processor:8082/execute-job"
	} else {
		host = "http://localhost:8082/execute-job"
	}
	go func() {
		defer wg.Done()
		resp, err := http.Post(host, "application/json", bytes.NewBuffer([]byte(jobFormated)))
		if err != nil {
			errChan <- err
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := ioutil.ReadAll(resp.Body)
			errChan <- errors.New(string(body) + "status" + strconv.Itoa(resp.StatusCode))
			return
		}
	}()
	go func() {
		wg.Wait()
		close(errChan)
	}()
	return <-errChan
}
