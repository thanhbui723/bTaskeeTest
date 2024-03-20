package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sending-service/models"
	"sending-service/package/util"
	"sending-service/services"

	"github.com/gin-gonic/gin"
)

type dtoAssignment struct {
	JobID    string `json:"job_id" validate:"required"`
	HelperID string `json:"helper_id" validate:"required"`
}

func CreateAssignment(c *gin.Context) {
	var request dtoAssignment
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

	// Get job information from booking service by REST API
	resp, err := http.Get("http://localhost:8072/api/jobs/" + request.JobID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Get job by id failed",
			"details": err.Error(),
		})
		return
	}
	defer resp.Body.Close()

	var job *models.Job
	if err := json.NewDecoder(resp.Body).Decode(&job); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Decode job failed",
			"details": err.Error(),
		})
		return
	}

	fmt.Println("job: ", job)

	// Get helper by helper_id to validate helper

	err = services.Assignment.CreateAssignment(request.JobID, request.HelperID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Create assignment failed",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, "success")
}
