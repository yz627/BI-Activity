package models

type College struct {
	ID                    uint64 `gorm:"primaryKey;autoIncrement"` // 主键，自动递增
	CollegeAccount        string `gorm:"type:varchar(64);unique;not null"` // 学院账号，唯一
	CollegeName           string `gorm:"type:varchar(64);not null"` // 学院名称
	Password              string `gorm:"type:varchar(255);not null"` // 密码
	AdminName             string `gorm:"type:varchar(255);"` // 管理员名称
	AdminIDNumber         string `gorm:"type:varchar(20);"` // 管理员身份证号
	AdminImageID          uint64 `gorm:"type:bigint"` // 管理员头像 ID
	AdminPhone            string `gorm:"type:varchar(64);unique;"` // 管理员电话，唯一
	AdminEmail            string `gorm:"type:varchar(255);unique;"` // 管理员邮箱，唯一
	Campus                int    `gorm:"type:int;not null"` // 校园 ID
	CollegeAddress        string `gorm:"type:varchar(255);not null"` // 学院地址
	CollegeIntroduction   string `gorm:"type:text"` // 学院简介
	CollegeAvatarID       uint64 `gorm:"type:bigint"` // 学院头像 ID
}