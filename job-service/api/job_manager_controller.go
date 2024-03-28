package api

import (
	"net/http"
	"strconv"

	"github.com/bk-rim/job-service/domain"
	"github.com/gin-gonic/gin"
)

type JobController struct {
	JobService domain.JobServicer
}

func NewJobController(js domain.JobServicer) *JobController {
	return &JobController{JobService: js}
}

func (jc *JobController) GetAll(c *gin.Context) {
	jobs, err := jc.JobService.GetAll()
	if err != nil {
		c.IndentedJSON(err.Status, gin.H{"message": err.Message})
		return
	}
	if jobs == nil {
		jobs = []domain.Job{}
	}
	c.IndentedJSON(http.StatusOK, jobs)
}

func (jc *JobController) GetByID(c *gin.Context) {
	id := c.Param("id")
	jobID, _ := strconv.Atoi(id)
	job, err := jc.JobService.GetByID(jobID)
	if err != nil {
		c.IndentedJSON(err.Status, gin.H{"message": err.Message})
		return
	}
	c.IndentedJSON(http.StatusOK, job)
}

func (jc *JobController) Create(broadcastMessage func(messageType string, data any)) gin.HandlerFunc {
	return func(c *gin.Context) {
		var jobPost domain.JobPost
		if err := c.BindJSON(&jobPost); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
			return
		}
		job, err := jc.JobService.Create(jobPost)
		if err != nil {
			c.IndentedJSON(err.Status, gin.H{"message": err.Message})
			return
		}
		broadcastMessage("job_created", job)
		c.IndentedJSON(http.StatusCreated, job)
	}
}

func (jc *JobController) Update(broadcastMessage func(messageType string, data any)) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		jobID, _ := strconv.Atoi(id)
		var jobUpdate domain.Job
		if err := c.BindJSON(&jobUpdate); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
			return
		}
		job, err := jc.JobService.Update(jobID, jobUpdate)
		if err != nil {
			c.IndentedJSON(err.Status, gin.H{"message": err.Message})
			return
		}
		broadcastMessage("job_updated", job)
		c.IndentedJSON(http.StatusOK, job)
	}
}

func (jc *JobController) Delete(c *gin.Context) {
	id := c.Param("id")
	jobID, _ := strconv.Atoi(id)
	err := jc.JobService.Delete(jobID)
	if err != nil {
		c.IndentedJSON(err.Status, gin.H{"message": err.Message})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Job deleted successfully"})
}
