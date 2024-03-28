package domain

type JobExecuterService struct {
	JobExecuterRepository JobExecuterRepositoryer
}

func NewJobExecuterService(jer JobExecuterRepositoryer) *JobExecuterService {
	return &JobExecuterService{JobExecuterRepository: jer}
}

func (jes *JobExecuterService) LaunchJobExecution(job Job) *ErrorJobService {
	err := jes.JobExecuterRepository.LaunchJobExecution(job)
	if err != nil {
		return &ErrorJobService{Message: err.Error(), Status: 500}
	}
	return nil
}
