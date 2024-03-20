package routers

import (
	"booking-service/routers/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/api/jobs", api.CreateJob)
	router.GET("/api/jobs/:id", api.GetJobByID)

	return router
}
