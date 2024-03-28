package api

import (
	"net/http"

	"github.com/bk-rim/job-service/domain"
	"github.com/gin-gonic/gin"
)

type JobExecuterController struct {
	JobExecuterService domain.JobExecuterServicer
}

func NewJobExecuterController(jes domain.JobExecuterServicer) *JobExecuterController {
	return &JobExecuterController{JobExecuterService: jes}
}

func (jec *JobExecuterController) LaunchJobExecution(c *gin.Context) {
	var job domain.Job
	if err := c.BindJSON(&job); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	errService := jec.JobExecuterService.LaunchJobExecution(job)
	if errService != nil {
		c.IndentedJSON(errService.Status, gin.H{"message": errService.Message})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Job execution launched successfully"})
}
