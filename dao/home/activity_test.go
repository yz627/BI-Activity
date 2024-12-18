package home

import (
	"bi-activity/configs"
	"bi-activity/dao"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestActivityDataCase_GetActivityListByID(t *testing.T) {
	conf := configs.InitConfig("./../configs/")
	data, fn := dao.NewDateDao(conf.Database, logrus.New())
	defer fn()

	activityDataCase := NewActivityDataCase(data, logrus.New())
	list, err := activityDataCase.GetActivityListByID(nil, []uint{1, 2, 3})
	if err != nil {
		t.Error(err)
	}
	for _, v := range list {
		t.Log(v)
	}
}
