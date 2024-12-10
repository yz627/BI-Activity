package main

import (
	"bi-activity/configs"
	"bi-activity/dao"
	"bi-activity/models"
	"github.com/sirupsen/logrus"
)

// 表的迁移
// TODO：命令行参数执行
func main() {
	conf := configs.InitConfig("./configs/")
	data, fn := dao.NewDateDao(conf.Database, logrus.New())
	defer fn()

	err := data.DB().AutoMigrate(
		&models.Student{},
		&models.Image{},
		&models.College{},
		&models.ActivityType{},
		&models.InviteCode{},
		&models.Activity{},
		&models.Admin{},
		&models.JoinCollegeAudit{},
		&models.Participant{},
		&models.StudentActivityAudit{},
		&models.CollegeRegistrationAudit{},
		&models.CollegeNameToAccount{},
		&models.Problem{},
	)
	if err != nil {
		panic(err)
	}
}
