package utils

import (
	"bi-activity/models"
	"bi-activity/service"
	"testing"
)

func TestStructCopy(t *testing.T) {
	img := &models.Image{
		FileName: "test",
		Url:      "test",
	}
	img.Model.ID = 1

	resImg := &service.RespImage{}

	StructCopy(img, resImg)
	t.Log(resImg)
}
