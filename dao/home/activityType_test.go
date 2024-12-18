package home

import (
	"bi-activity/configs"
	"bi-activity/dao"
	"context"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestActivityTypeDataCase_GetActivityAllTypes(t *testing.T) {
	conf := configs.InitConfig("./../configs/")
	data, fn := dao.NewDateDao(conf.Database, logrus.New())
	defer fn()

	activityTypeDataCase := NewActivityTypeDataCase(data, logrus.New())
	list, err := activityTypeDataCase.GetActivityAllTypes(context.Background())
	if err != nil {
		t.Error(err)
	}
	for _, v := range list {
		t.Log(v, v.Image)
	}
}
