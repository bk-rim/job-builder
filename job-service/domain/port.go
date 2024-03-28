package domain

type JobRepositoryer interface {
	GetAll() ([]Job, error)
	GetByID(id int) (Job, error)
	Create(job *Job) (Job, error)
	Update(id int, job Job) (Job, error)
	Delete(id int) error
}

type JobExecuterRepositoryer interface {
	LaunchJobExecution(job Job) error
}

type JobServicer interface {
	GetAll() ([]Job, *ErrorJobService)
	GetByID(id int) (Job, *ErrorJobService)
	Create(jobPost JobPost) (Job, *ErrorJobService)
	Update(id int, job Job) (Job, *ErrorJobService)
	Delete(id int) *ErrorJobService
}

type JobExecuterServicer interface {
	LaunchJobExecution(job Job) *ErrorJobService
}
