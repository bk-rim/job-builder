package domain_utils

import (
	"time"

	"github.com/bk-rim/job-processor/domain"
)

func IsReadyToBeProcessed(job domain.Job) bool {
	switch job.Frequency {
	case "daily":
		if job.ExecutedOn.IsZero() {
			return job.CreatedOn.Add(24 * time.Hour).Before(time.Now())
		} else {
			return job.ExecutedOn.Add(24 * time.Hour).Before(time.Now())
		}

	case "weekly":
		if job.ExecutedOn.IsZero() {
			return job.CreatedOn.Truncate(24 * time.Hour).Add(7 * 24 * time.Hour).Before(time.Now().Truncate(24 * time.Hour))
		} else {
			return job.ExecutedOn.Truncate(24 * time.Hour).Add(7 * 24 * time.Hour).Before(time.Now().Truncate(24 * time.Hour))
		}
	default:
		return false
	}
}
