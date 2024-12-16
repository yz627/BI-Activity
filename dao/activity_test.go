package dao

import (
	"bi-activity/configs"
	"context"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestActivityDataCase_GetActivityAllTypes(t *testing.T) {
	conf := configs.InitConfig("./../configs/")
	data, fn := NewDateDao(conf.Database, logrus.New())
	defer fn()

	activityData := NewActivityDataCase(data, logrus.New())
	list, err := activityData.GetActivityAllTypes(context.TODO())
	if err != nil {
		t.Error(err)
	}
	t.Log(list[0])
}
