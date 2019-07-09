package api

import (
	"github.com/gin-gonic/gin"
	"Gotube/service"
)

// 创建视频接口
func CreateVideo(c *gin.Context) {
	service := service.CreateVideoService {}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}