package domain_test

import (
	"errors"
	"net/http"
	"strconv"
	"testing"

	"github.com/bk-rim/job-service/domain"
	"github.com/stretchr/testify/assert"
)

type MockJobExecuterRepository struct{}

type RequestSender struct {
	StatusCode int
}

func (s *RequestSender) Post(url string) (*http.Response, error) {
	return &http.Response{StatusCode: s.StatusCode, Body: http.NoBody}, nil
}

func (r *MockJobExecuterRepository) LaunchJobExecution(job domain.Job) error {
	requestSender := &RequestSender{StatusCode: http.StatusOK}
	resp, err := requestSender.Post("http://localhost:8082/execute-job")
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("request error with status" + strconv.Itoa(resp.StatusCode))
	}
	return nil
}

func TestLaunchJobExecution(t *testing.T) {
	jes := domain.NewJobExecuterService(&MockJobExecuterRepository{})
	err := jes.LaunchJobExecution(domain.Job{})
	assert.Nil(t, err)
}
