package domain_utils

import (
	"testing"
	"time"

	"github.com/bk-rim/job-processor/domain"
)

func TestIsReadyToBeProcessed(t *testing.T) {
	dailyJob := domain.Job{
		Frequency:  "daily",
		CreatedOn:  time.Now().Add(-48 * time.Hour),
		ExecutedOn: time.Time{},
	}
	expectedDaily := true
	if result := IsReadyToBeProcessed(dailyJob); result != expectedDaily {
		t.Errorf("Expected %v, but got %v for daily job", expectedDaily, result)
	}

	weeklyJob := domain.Job{
		Frequency:  "weekly",
		CreatedOn:  time.Now().Add(-14 * 24 * time.Hour),
		ExecutedOn: time.Time{},
	}
	expectedWeekly := true
	if result := IsReadyToBeProcessed(weeklyJob); result != expectedWeekly {
		t.Errorf("Expected %v, but got %v for weekly job", expectedWeekly, result)
	}

	defaultJob := domain.Job{
		Frequency:  "monthly",
		CreatedOn:  time.Now().Add(-30 * 24 * time.Hour),
		ExecutedOn: time.Time{},
	}
	expectedDefault := false
	if result := IsReadyToBeProcessed(defaultJob); result != expectedDefault {
		t.Errorf("Expected %v, but got %v for default job", expectedDefault, result)
	}
}
