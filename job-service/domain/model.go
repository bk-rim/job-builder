package domain

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Job struct {
	ID           int        `json:"id"`
	Name         string     `json:"name"`
	Type         string     `json:"type"`
	Frequency    string     `json:"frequency"`
	CreatedOn    time.Time  `json:"created_on"`
	Status       string     `json:"status"`
	ExecutedOn   *time.Time `json:"executed_on"`
	WebhookSlack string     `json:"webhook_slack"`
}

type JobPost struct {
	Name         string `json:"name" validate:"required"`
	Type         string `json:"type" validate:"required"`
	Frequency    string `json:"frequency" validate:"required"`
	WebhookSlack string `json:"webhook_slack" validate:"required"`
}

type ErrorJobService struct {
	Message string
	Status  int
}

func (jp *JobPost) Validate() error {
	validate := validator.New()
	return validate.Struct(jp)
}
