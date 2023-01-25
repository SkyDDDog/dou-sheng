package repository

type VideoComment struct {
	CommentId   int64 `gorm:"primarykey"`
	VideoId     int64
	AuthorId    int64
	CommentText string `gorm:"type:longtext"`
	Model
}
