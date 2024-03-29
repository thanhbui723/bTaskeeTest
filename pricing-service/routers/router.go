package routers

import (
	"pricing-service/routers/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/api/prices", api.GetPrice)

	return router
}
