package routers

import (
	"sending-service/routers/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/api/assignments", api.CreateAssignment)
	router.GET("/api/helpers", api.GetHelpers)
	return router
}
