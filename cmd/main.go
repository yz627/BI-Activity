package main

import (
	"bi-activity/configs"
	"bi-activity/dao"

	"github.com/sirupsen/logrus"
)

func main() {
	config := configs.InitConfig()

	logger := logrus.New()

	dao.NewDataDao(config.Database, logger)
}
