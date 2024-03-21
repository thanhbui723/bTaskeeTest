package api

import (
	"errors"
	"net/http"
	"sending-service/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// API Create Helper

// API Get Helper by HelperID

// API Get helpers
func GetHelpers(c *gin.Context) {
	helpers, err := models.Repository.Helper.GetHelpers()
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, gin.H{
				"error":   "Helpers not found",
				"details": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Get list helper failed",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"helpers": helpers,
	})
}
