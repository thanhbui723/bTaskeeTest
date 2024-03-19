package api

import (
	"net/http"
	"pricing-service/services"
	"time"

	"github.com/gin-gonic/gin"
)

func GetPrice(c *gin.Context) {
	dateStr := c.Param("date")
	jobType := c.Query("type") // Lấy tham số loại công việc từ query string

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Lỗi: Định dạng ngày không hợp lệ"})
		return
	}

	price, err := services.Price.GetPriceByDateAndType(date, jobType)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"price": price.Price})
}
