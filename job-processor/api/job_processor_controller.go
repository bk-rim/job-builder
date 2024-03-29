package api

import (
	"net/http"

	"github.com/bk-rim/job-processor/domain"
	"github.com/gin-gonic/gin"
)

type JobProcessorController struct {
	JobProcessorService domain.JobProcessorServicer
}

func NewJobProcessorController(jps domain.JobProcessorServicer) *JobProcessorController {
	return &JobProcessorController{JobProcessorService: jps}
}

func (jpc *JobProcessorController) ProcessJob(c *gin.Context) {
	var job domain.Job
	if err := c.BindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := jpc.JobProcessorService.ProcessJob(&job)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "job processed successfully"})
}
