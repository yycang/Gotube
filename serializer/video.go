package serializer

import "Gotube/model"


// Video 视频序列化器
type Video struct {
	Id 		uint 	`json:"id"`
	Title	string	`json:"title"`
	Info	string 	`json:"info"`
	CreatedAt	int64 	`json:"created_at"`
}

// BuildVideo 序列化视频
func BuildVideo(item model.Video) Video {
	return Video{
		Id:		item.ID,
		Title:	item.Title,
		Info:	item.Info,
		CreatedAt:	item.CreatedAt.Unix(),
	}
}