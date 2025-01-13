package main

import (
	"bi-activity/configs"
	"bi-activity/dao"
	"bi-activity/models"
	"github.com/sirupsen/logrus"
)

// 表的迁移
func main() {

	conf := configs.InitConfig("../../configs/")

	data, fn := dao.NewDataDao(conf.Database, logrus.New())
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
		&models.Message{},
		&models.Conversation{},
	)
	if err != nil {
		panic(err)
	}
}
