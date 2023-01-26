package repository

import "gorm.io/gorm"

// BeforeCreate	创建前Hook，生成雪花算法分布式主键
func (uf *UserFollow) BeforeCreate(tx *gorm.DB) (err error) {
	uf.Id = int64(GenID())
	return nil
}
