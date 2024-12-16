package dao

import (
	"bi-activity/configs"
	"bi-activity/response"
	"bi-activity/response/errors"
	"bi-activity/service"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestImageDataCase_GetImageByType(t *testing.T) {
	conf := configs.InitConfig("./../configs/")
	data, fn := NewDateDao(conf.Database, logrus.New())
	defer fn()

	imageDataCase := NewImageDataCase(data, logrus.New())
	list, err := imageDataCase.GetImageByType(nil, 1)
	if err != nil {
		t.Error(response.Failf(errors.ServerError, "获取图片失败"))
	}
	_, resp := response.Success(list)
	t.Log(resp.Data.([]*service.Image)[0])
}
