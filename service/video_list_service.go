package service

import (
	"Gotube/model"
	"Gotube/serializer"
)

type VideoListService struct {
}

func (service *VideoListService) List() serializer.Response {
	var videos []model.Video
	err := model.DB.Find(&videos).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg: "数据库链接错误",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildVideos(videos),
	}
}