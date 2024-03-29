package main

import (
	adapter "github.com/bk-rim/job-processor/adapters"
	"github.com/bk-rim/job-processor/adapters/repository"
	"github.com/bk-rim/job-processor/api"
	service "github.com/bk-rim/job-processor/domain/services"
	"github.com/gin-gonic/gin"
)

func main() {

	jobProcessorRepo := repository.NewJobProcessorRepository()
	weatherAPI := adapter.NewWeatherAPIAdapter()
	bridgeStatusAPI := adapter.NewBridgeStatusAPIAdapter()
	jobProcessorService := service.NewJobProcessorService(weatherAPI, bridgeStatusAPI, jobProcessorRepo)
	jobProcessorController := api.NewJobProcessorController(jobProcessorService)

	go jobProcessorService.StartProcessingJobs()

	router := gin.Default()
	router.POST("/execute-job", jobProcessorController.ProcessJob)

	router.Run(":8082")
}
