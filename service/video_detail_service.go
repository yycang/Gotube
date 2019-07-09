package service

import (
	"Gotube/model"
	"Gotube/serializer"
)

type VideoDetailService struct {
	Title string `form:"title" json:"title" bindling:"required,min=2,max=30"`
	Info string `form:"info" json:"info" binding:"max=300"`
}

func (service *VideoDetailService) Detail(id string) serializer.Response {
	video := model.Video{}
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Status: 40004,
			Msg: "视频不存在",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}