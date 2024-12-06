package dao

import (
	"bi-activity/configs"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Data struct {
	log *logrus.Logger
	db  *gorm.DB
}

type Redis struct {
	rdb redis.Cmdable
}

func NewDateDao(conf *configs.Database, log *logrus.Logger) *Data {
	return &Data{}
}

func NewRedisDao(conf *configs.Redis, log *logrus.Logger) *Redis {
	return &Redis{}
}
