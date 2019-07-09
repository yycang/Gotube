package service

import (
	"Gotube/model"
	"Gotube/serializer"
)

type CreateVideoService struct {
	Title string `form:"title" json:"title" bindling:"required,min=2,max=30"`
	Info string `form:"info" json:"info" binding:"max=300"`
}

func (service *CreateVideoService) Create() serializer.Response {
	video := model.Video{
		Title: service.Title,
		Info: service.Info,
	}
	err := model.DB.Create(&video).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg: "视频保存失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}