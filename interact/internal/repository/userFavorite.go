package repository

type UserFavorite struct {
	Id      int64 `gorm:"primarykey"`
	UserId  int64
	VideoId int64
	Model
}
