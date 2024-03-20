package api

import (
	"fmt"
	"net/http"
	"pricing-service/package/util"
	"pricing-service/services"
	"time"

	"github.com/gin-gonic/gin"
)

type dtoPrice struct {
	Date     int            `json:"date" form:"date" validate:"required"`
	JobType  util.JobType   `json:"job_type" form:"job_type" validate:"required"`
	DateType util.PriceType `json:"date_type" form:"date_type" validate:"required"`
	Duration int            `json:"duration" form:"duration" validate:"required"` // hours
}

func GetPrice(c *gin.Context) {
	var request dtoPrice
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, "Bind query param failed: "+err.Error())
		return
	}
	fmt.Println("Request input: ", request)

	// Validate request input
	if err := util.Validator.Struct(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Missing query param",
			"details": err.Error(),
		})
		return
	}

	// convert time Unix to time.Time and round to day
	dateTime := time.Unix(int64(request.Date), 0)
	dateTime = util.RoundToDay(dateTime)

	price, err := services.Price.GetPriceByDateAndType(dateTime, string(request.JobType), string(request.DateType), request.Duration)
	if err != nil {
		if err == services.ErrPriceNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error":   "Price not upload",
				"details": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Get price failed",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"price": price,
	})
}
