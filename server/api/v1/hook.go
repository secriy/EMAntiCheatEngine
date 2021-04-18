package v1

import (
	"github.com/gin-gonic/gin"
	"server/service/hooklog"
)

// CountHookLog
func CountHookLog(c *gin.Context) {
	countHookLogService := hooklog.CountHookExceptionService{}
	res := countHookLogService.Count()
	c.JSON(200, res)
}

// ListHookLog
func ListHookLog(c *gin.Context) {
	listHookLogService := hooklog.ListHookLogService{}
	if err := c.ShouldBind(&listHookLogService); err == nil {
		res := listHookLogService.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ListExpHookLog
func ListExpHookLog(c *gin.Context) {
	listHookLogService := hooklog.ListHookLogService{}
	if err := c.ShouldBind(&listHookLogService); err == nil {
		res := listHookLogService.ListExp()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
