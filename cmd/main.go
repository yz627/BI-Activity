package main

import (
	"bi-activity/configs"
	"bi-activity/dao"
	"bi-activity/router"
	"bi-activity/testcase"

	"github.com/sirupsen/logrus"
)

func main() {
	config := configs.InitConfig()

	logger := logrus.New()

	data := dao.NewDataDao(config.Database, logger)
	
	testcase.RunStudentTest(data.Db)

	r := router.InitRouter()
	
	r.Run(":8080")
}
