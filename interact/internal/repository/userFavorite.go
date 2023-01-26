package repository

import (
	"interact/internal/service"
)

type UserFavorite struct {
	Id      int64 `gorm:"primarykey"`
	UserId  int64
	VideoId int64
	Model
}

func (*UserFavorite) CreateUf(req *service.DouyinFavoriteActionRequest) error {
	uf := UserFavorite{
		UserId:  req.UserId,
		VideoId: req.VideoId,
		Model:   Model{},
	}
	return DB.Create(&uf).Error
}

func (*UserFavorite) DeleteUf(req *service.DouyinFavoriteActionRequest) error {
	return DB.Unscoped().Where("video_id = ? and user_id = ?", req.VideoId, req.UserId).Delete(&UserFavorite{}).Error
}

func (*UserFavorite) FavoriteList(req *service.DouyinFavoriteListRequest) (ufList []UserFavorite, err error) {
	err = DB.Model(&UserFavorite{}).
		Where("user_id = ?", req.UserId).
		Order("create_date desc").
		Find(&ufList).Error
	if err != nil {
		return nil, err
	}
	return ufList, nil
}

func BuildVideo(item UserFavorite) *service.Video {
	return &service.Video{
		Id: item.VideoId,
		Author: &service.User{
			Id: item.Id,
		},
		IsFavorite: true,
	}
}

func BuildVideos(item []UserFavorite) (videoList []*service.Video) {
	for _, v := range item {
		video := BuildVideo(v)
		videoList = append(videoList, video)
	}
	return videoList
}

func (*UserFavorite) FavoriteCount(videoId int64) (count int64) {
	DB.Model(&UserFavorite{}).Where("video_id = ?", videoId).Count(&count)
	return count
}

func (*UserFavorite) IsFavorite(videoId int64, userId int64) bool {
	var count int64
	err := DB.Model(&UserFavorite{}).
		Where("video_id = ? and user_id = ?", videoId, userId).
		Count(&count).Error
	if err != nil {
		return false
	}
	return 0 < count
}
