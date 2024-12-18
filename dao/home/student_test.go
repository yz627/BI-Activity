package home

import (
	"bi-activity/configs"
	"bi-activity/dao"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestStudentDataCase_GetCollegeStudentCount(t *testing.T) {
	conf := configs.InitConfig("./../../configs/")
	data, fn := dao.NewDateDao(conf.Database, logrus.New())
	defer fn()

	studentData := NewStudentDataCase(data, logrus.New())
	total, err := studentData.GetStudentTotal(nil)
	if err != nil {
		t.Error(err)
	}
	t.Log(total)

	count, err := studentData.GetCollegeStudentCount(nil)
	if err != nil {
		t.Error(err)
	}

	for college, num := range count {
		t.Log(college, num)
	}
}
