package repository

import "gorm.io/gorm"

// BeforeCreate	创建前Hook，生成雪花算法分布式主键
func (v *Video) BeforeCreate(tx *gorm.DB) (err error) {
	if v.VideoId != 0 {
		v.VideoId = int64(GenID())
	}
	return nil
}
