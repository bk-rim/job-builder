package domain_test

import (
	"testing"
	"time"

	"github.com/bk-rim/job-service/domain"
	"github.com/stretchr/testify/assert"
)

var dataJob = []domain.Job{
	{
		ID:           1,
		Name:         "job1",
		Type:         "weather",
		Frequency:    "daily",
		CreatedOn:    time.Now(),
		Status:       "pending",
		ExecutedOn:   nil,
		WebhookSlack: "https://slack.com/webhook",
	},
	{
		ID:           2,
		Name:         "job2",
		Type:         "bridge_status",
		Frequency:    "weekly",
		CreatedOn:    time.Now(),
		Status:       "pending",
		ExecutedOn:   nil,
		WebhookSlack: "https://slack.com/webhook",
	},
}

type JobRepositoryMock struct{}

func (jrm *JobRepositoryMock) GetAll() ([]domain.Job, error) {

	return dataJob, nil
}

func (jrm *JobRepositoryMock) GetByID(id int) (domain.Job, error) {
	for _, job := range dataJob {
		if job.ID == id {
			return job, nil
		}
	}
	return domain.Job{}, nil
}

func (jrm *JobRepositoryMock) Create(job *domain.Job) (domain.Job, error) {
	job.ID = 3
	dataJob = append(dataJob, *job)
	return *job, nil
}

func (jrm *JobRepositoryMock) Update(id int, job domain.Job) (domain.Job, error) {
	for i, j := range dataJob {
		if j.ID == id {
			dataJob[i] = job
			return job, nil
		}
	}
	return domain.Job{}, nil
}

func (jrm *JobRepositoryMock) Delete(id int) error {
	for i, j := range dataJob {
		if j.ID == id {
			dataJob = append(dataJob[:i], dataJob[i+1:]...)
			return nil
		}
	}
	return nil
}

func TestGetAll(t *testing.T) {
	jobService := domain.NewJobService(&JobRepositoryMock{})
	jobs, err := jobService.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, dataJob, jobs)
}

func TestGetByID(t *testing.T) {
	jobService := domain.NewJobService(&JobRepositoryMock{})
	job, err := jobService.GetByID(1)
	assert.Nil(t, err)
	assert.Equal(t, dataJob[0], job)
}

func TestCreate(t *testing.T) {
	jobService := domain.NewJobService(&JobRepositoryMock{})
	jobPost := domain.JobPost{
		Name:         "job3",
		Type:         "weather",
		Frequency:    "daily",
		WebhookSlack: "https://slack.com/webhook",
	}
	job, err := jobService.Create(jobPost)
	assert.Nil(t, err)
	index := len(dataJob) - 1
	assert.Equal(t, dataJob[index].ID, job.ID)
}

func TestUpdate(t *testing.T) {
	jobService := domain.NewJobService(&JobRepositoryMock{})
	job := domain.Job{
		ID:           1,
		Name:         "job1",
		Type:         "weather",
		Frequency:    "daily",
		CreatedOn:    time.Now(),
		Status:       "completed",
		ExecutedOn:   nil,
		WebhookSlack: "https://slack.com/webhook",
	}
	jobUpdated, err := jobService.Update(1, job)
	assert.Nil(t, err)
	assert.Equal(t, job, jobUpdated)
}

func TestDelete(t *testing.T) {
	jobService := domain.NewJobService(&JobRepositoryMock{})
	err := jobService.Delete(1)
	assert.Nil(t, err)
	_, err = jobService.GetByID(1)
	assert.NotNil(t, err)
}
