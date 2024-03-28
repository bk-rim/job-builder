package domain

type JobProcessorRepositoryer interface {
	UpadateJobStatus(job *Job) error
}

type JobProcessorServicer interface {
	ProcessJob(job *Job) error
}

type ExternalWeatherAPI interface {
	FetchWeatherData() (string, error)
}

type ExternalBridgeStatusAPI interface {
	GetBridgeData() (string, error)
	ProcessBridgeStatus(data string) (string, error)
}
