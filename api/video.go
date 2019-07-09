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

// 查看视频接口
func VideoDetail(c *gin.Context)  {
	service := service.VideoDetailService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Detail(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 视频列表接口
func VideoList(c *gin.Context)  {
	service := service.VideoListService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}