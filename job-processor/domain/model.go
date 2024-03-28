package domain

import "time"

type Job struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Type         string    `json:"type"`
	Frequency    string    `json:"frequency"`
	CreatedOn    time.Time `json:"created_on"`
	Status       string    `json:"status"`
	ExecutedOn   time.Time `json:"executed_on"`
	WebhookSlack string    `json:"webhook_slack"`
}
