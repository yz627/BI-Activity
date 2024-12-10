package dao

import (
	"bi-activity/configs"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestNewDateDao(t *testing.T) {
	conf := configs.InitConfig("./../configs/")
	db, fn := NewDateDao(conf.Database, logrus.New())
	defer fn()

	t.Log(db)
}

func TestNewRedisDao(t *testing.T) {
	conf := configs.InitConfig("./../configs/")
	redis, fn := NewRedisDao(conf.Redis, logrus.New())
	defer fn()
	t.Log(redis)
}
