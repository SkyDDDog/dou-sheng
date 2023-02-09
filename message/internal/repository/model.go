package repository

import "time"
import "gorm.io/plugin/soft_delete"

type Model struct {
	Del_flag    soft_delete.DeletedAt `gorm:"comment:逻辑删除;softDelete:flag;type:char"`
	Create_date time.Time             `gorm:"comment:创建时间;autoCreateTime"`
	Update_date time.Time             `gorm:"comment:更新时间;autoUpdateTime"`
	Remarks     string                `gorm:"comment:备注;type:varchar(255)"`
}
