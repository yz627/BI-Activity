package dao

import (
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
