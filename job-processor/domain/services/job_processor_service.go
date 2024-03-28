package service

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/bk-rim/job-processor/domain"
	domain_utils "github.com/bk-rim/job-processor/domain/utils"
)

type JobProcessorService struct {
	WeatherAPI             domain.ExternalWeatherAPI
	BridgeStatusAPI        domain.ExternalBridgeStatusAPI
	JobProcessorRepository domain.JobProcessorRepositoryer
}

func NewJobProcessorService(weatherAPI domain.ExternalWeatherAPI, bridgeStatusAPI domain.ExternalBridgeStatusAPI, jpr domain.JobProcessorRepositoryer) *JobProcessorService {
	return &JobProcessorService{
		WeatherAPI:             weatherAPI,
		BridgeStatusAPI:        bridgeStatusAPI,
		JobProcessorRepository: jpr,
	}
}

func (s *JobProcessorService) ProcessJob(job *domain.Job) error {

	job.ExecutedOn = time.Now()
	err := s.JobProcessorRepository.UpadateJobStatus(job)
	if err != nil {
		return err
	}

	switch job.Type {
	case "weather":
		job.Status = "on going"
		err = s.JobProcessorRepository.UpadateJobStatus(job)
		if err != nil {
			return err
		}
		weatherData, err := s.WeatherAPI.FetchWeatherData()
		if err != nil {
			job.Status = "failed"
			s.JobProcessorRepository.UpadateJobStatus(job)
			return err
		}

		statusCode, err := domain_utils.PostOnWebhookSlack(weatherData, *job)
		if err != nil {
			job.Status = "failed"
			s.JobProcessorRepository.UpadateJobStatus(job)
			return err
		}
		log.Printf("Weather data posted on Slack with status code: %s\n", statusCode)

	case "bridge_status":
		job.Status = "on going"
		err = s.JobProcessorRepository.UpadateJobStatus(job)
		if err != nil {
			return err
		}
		statusData, err := s.BridgeStatusAPI.GetBridgeData()
		if err != nil {
			job.Status = "failed"
			s.JobProcessorRepository.UpadateJobStatus(job)

			return err

		}
		statusDataProcessed, err := s.BridgeStatusAPI.ProcessBridgeStatus(statusData)
		if err != nil {
			job.Status = "failed"
			s.JobProcessorRepository.UpadateJobStatus(job)
			return err
		}

		statusCode, err := domain_utils.PostOnWebhookSlack(statusDataProcessed, *job)
		if err != nil {
			job.Status = "failed"
			s.JobProcessorRepository.UpadateJobStatus(job)
			return err
		}
		log.Printf("Bridge status posted on Slack with status code: %s\n", statusCode)

	default:
		return errors.New("unknown job type")

	}

	job.Status = "executed"
	log.Printf("Job %d executed successfully\n", job.ID)
	err = s.JobProcessorRepository.UpadateJobStatus(job)
	if err != nil {
		return err
	}

	return nil
}

func (s *JobProcessorService) StartProcessingJobs() {
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			jobs, err := domain_utils.FetchJobs()
			if err != nil {
				fmt.Println("Error fetching jobs:", err)
				continue
			}

			for _, job := range jobs {

				if domain_utils.IsReadyToBeProcessed(job) {

					err := s.ProcessJob(&job)
					if err != nil {
						log.Printf("Error processing job %d: %v\n", job.ID, err)
					}
				}
			}
		}
	}
}
