package models

type Image struct {
	ID       uint64 `gorm:"primaryKey;autoIncrement"` // 主键，自动递增
	FileName string `gorm:"type:varchar(255);not null"` // 文件名
	URL      string `gorm:"type:varchar(255);not null"` // 文件 URL
	Type     int    `gorm:"type:int;not null"` // 文件类型
}