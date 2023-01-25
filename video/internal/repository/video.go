package repository

import (
	"github.com/spf13/viper"
	"strconv"
	"video/internal/service"
)

type Video struct {
	VideoId  int64 `gorm:"primarykey"`
	AuthorId int64
	Title    string
	Model
}

// CreateVideo	创建新视频
func (*Video) CreateVideo(req *service.DouyinPublishActionRequest) (video Video, err error) {
	video = Video{
		AuthorId: req.UserId,
		Title:    req.Title,
		Model:    Model{},
	}
	err = DB.Create(&video).Error
	return video, err
}

// VideoShow	发布列表
func (*Video) VideoShow(req *service.DouyinPublishListRequest) (videoList []Video, err error) {
	err = DB.Model(&Video{}).Where("author_id = ?", req.UserId).Order("create_date desc").Find(&videoList).Error
	if err != nil {
		return nil, err
	}
	return videoList, err
}

// VideoFeed	视频流接口
func (*Video) VideoFeed(req *service.DouyinFeedRequest) (videoList []Video, err error) {
	err = DB.Model(&Video{}).Where("create_date > ?", req.LatestTime).
		Order("create_date asc").Limit(10).Find(&videoList).Error
	if err != nil {
		return nil, err
	}
	return videoList, err
}

func BuildVideo(item Video) *service.Video {
	prefix := "http://" + viper.GetString("server.host") + ":9999/static/"
	return &service.Video{
		Id:            item.VideoId,
		Author:        nil,
		PlayUrl:       prefix + "video/" + strconv.FormatInt(item.VideoId, 10) + ".mp4",
		CoverUrl:      prefix + "cover/" + strconv.FormatInt(item.VideoId, 10) + ".png",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         item.Title,
	}
}

func BuildVideos(item []Video) (vList []*service.Video) {
	for _, v := range item {
		video := BuildVideo(v)
		vList = append(vList, video)
	}
	return vList
}
