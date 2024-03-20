package api

import (
	"booking-service/models"
	"booking-service/package/util"
	"booking-service/services"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type dtoJob struct {
	Date        int          `json:"date" validate:"required"`
	Type        util.JobType `json:"type"  validate:"required"`
	Description string       `json:"description"`
	Duration    int          `json:"duration" validate:"required"`
	Price       int          `json:"price" validate:"required"`
	Address     string       `json:"address" validate:"required"`
}

func CreateJob(c *gin.Context) {
	var request dtoJob
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, "Parse JSON body failed: "+err.Error())
		return
	}
	fmt.Println("Request input: ", request)

	// Validate request input
	if err := util.Validator.Struct(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Validate JSON body failed",
			"details": err.Error(),
		})
		return
	}

	// convert time Unix to time.Time and round to day
	dateTime := time.Unix(int64(request.Date), 0)
	dateTime = util.RoundToDay(dateTime)

	job := &models.Job{
		Date:        dateTime,
		Description: request.Description,
		Type:        request.Type,
		Status:      util.Pending,
		Price:       request.Price,
		Address:     request.Address,
		Duration:    request.Duration,
	}

	newJob, err := models.Repository.Job.CreateJob(job)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Create job failed",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, newJob)
}

func GetJobByID(c *gin.Context) {
	jobID := c.Param("id")

	fmt.Println("jobID: ", jobID)
	job, err := services.Job.GetJobByID(jobID)
	if err != nil {
		if err == services.ErrJobNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error":   "Job not found",
				"details": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Get job failed",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, job)
}
