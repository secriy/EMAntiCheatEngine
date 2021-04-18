package v1

import (
	"github.com/gin-gonic/gin"
	"server/service/antilog"
)

// CountAntiLog
func CountAntiLog(c *gin.Context) {
	countAntiLogService := antilog.CountAntiExceptionService{}
	res := countAntiLogService.Count()
	c.JSON(200, res)
}

// ListAntiLog
func ListAntiLog(c *gin.Context) {
	listAntiLogService := antilog.ListAntiLogService{}
	if err := c.ShouldBind(&listAntiLogService); err == nil {
		res := listAntiLogService.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ListExpAntiLog
func ListExpAntiLog(c *gin.Context) {
	listAntiLogService := antilog.ListAntiLogService{}
	if err := c.ShouldBind(&listAntiLogService); err == nil {
		res := listAntiLogService.ListExp()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
