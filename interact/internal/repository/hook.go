package repository

import "gorm.io/gorm"

// BeforeCreate	创建前Hook，生成雪花算法分布式主键
func (uf *UserFavorite) BeforeCreate(tx *gorm.DB) (err error) {
	uf.Id = int64(GenID())
	return nil
}

// BeforeCreate	创建前Hook，生成雪花算法分布式主键
func (comment *VideoComment) BeforeCreate(tx *gorm.DB) (err error) {
	comment.CommentId = int64(GenID())
	return nil
}
