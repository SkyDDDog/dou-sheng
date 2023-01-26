package repository

import "interact/internal/service"

type VideoComment struct {
	CommentId   int64 `gorm:"primarykey"`
	VideoId     int64
	AuthorId    int64
	CommentText string `gorm:"type:longtext"`
	Model
}

func (*VideoComment) CreateComment(req *service.DouyinCommentActionRequest) (vc VideoComment, err error) {
	vc = VideoComment{
		VideoId:     req.VideoId,
		AuthorId:    req.UserId,
		CommentText: req.CommentText,
		Model:       Model{},
	}
	err = DB.Create(&vc).Error
	return vc, err
}

func (*VideoComment) DeleteComment(req *service.DouyinCommentActionRequest) (vc VideoComment, err error) {
	err = DB.Where("comment_id = ?", req.CommentId).Delete(&vc).Error
	return vc, err
}

func BuildVideoComment(item VideoComment) *service.Comment {
	return &service.Comment{
		Id:         item.CommentId,
		User:       &service.User{Id: item.AuthorId},
		Content:    item.CommentText,
		CreateDate: item.Model.Create_date.Format("01-02"),
	}
}

func BuildVideoComments(item []VideoComment) (vcList []*service.Comment) {
	for _, videoComment := range item {
		vc := BuildVideoComment(videoComment)
		vcList = append(vcList, vc)
	}
	return vcList
}

func (*VideoComment) CommentList(req *service.DouyinCommentListRequest) (vcList []VideoComment, err error) {
	err = DB.Model(&VideoComment{}).
		Where("video_id = ?", req.VideoId).
		Find(&vcList).Error
	return vcList, err
}

func (*VideoComment) CommentCount(videoId int64) (count int64) {
	DB.Model(&VideoComment{}).Where("video_id = ?", videoId).Count(&count)
	return count
}
