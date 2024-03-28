package domain

import "time"

type JobService struct {
	JobRepository JobRepositoryer
}

func NewJobService(jr JobRepositoryer) *JobService {
	return &JobService{JobRepository: jr}
}

func (js *JobService) GetAll() ([]Job, *ErrorJobService) {
	jobs, err := js.JobRepository.GetAll()
	if err != nil {
		return nil, &ErrorJobService{Message: err.Error(), Status: 500}
	}
	return jobs, nil
}

func (js *JobService) GetByID(id int) (Job, *ErrorJobService) {
	job, err := js.JobRepository.GetByID(id)
	if err != nil {
		return Job{}, &ErrorJobService{Message: err.Error(), Status: 500}
	}
	if job.ID == 0 {
		return Job{}, &ErrorJobService{Message: "Job not found", Status: 404}
	}
	return job, nil
}

func (js *JobService) Create(jobPost JobPost) (Job, *ErrorJobService) {

	err := jobPost.Validate()
	if err != nil {
		return Job{}, &ErrorJobService{Message: err.Error(), Status: 400}
	}
	job := Job{
		Name:         jobPost.Name,
		Type:         jobPost.Type,
		Frequency:    jobPost.Frequency,
		CreatedOn:    time.Now(),
		Status:       "pending",
		WebhookSlack: jobPost.WebhookSlack,
	}

	if jobPost.Type != "weather" && jobPost.Type != "bridge_status" {
		return Job{}, &ErrorJobService{Message: "Invalid job type", Status: 400}
	}

	if jobPost.Frequency != "daily" && jobPost.Frequency != "weekly" {
		return Job{}, &ErrorJobService{Message: "Invalid job frequency", Status: 400}
	}

	jobCreated, err := js.JobRepository.Create(&job)
	if err != nil {
		return Job{}, &ErrorJobService{Message: err.Error(), Status: 500}
	}
	return jobCreated, nil
}

func (js *JobService) Update(id int, job Job) (Job, *ErrorJobService) {
	job, err := js.JobRepository.Update(id, job)
	if err != nil {
		return Job{}, &ErrorJobService{Message: err.Error(), Status: 500}
	}
	return job, nil
}

func (js *JobService) Delete(id int) *ErrorJobService {
	err := js.JobRepository.Delete(id)
	if err != nil {
		return &ErrorJobService{Message: err.Error(), Status: 500}
	}
	return nil
}
