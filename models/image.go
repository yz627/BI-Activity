package models

type Image struct {
	ID       uint64 `gorm:"primaryKey;autoIncrement" json:"id"`          // 主键，自动递增
	FileName string `gorm:"type:varchar(255);not null" json:"file_name"` // 文件名
	URL      string `gorm:"type:varchar(255);not null" json:"url"`       // 文件 URL
	Type     int    `gorm:"type:int;not null" json:"type"`               // 文件类型
}
