package utils

import (
	"testing"
)

type User struct {
	ID    uint
	Name  string
	Image *Image
}

type Image struct {
	ID  uint
	Url string
}

type UserInfo struct {
	ID   uint
	Name string
	Url  string
}

func TestStructCopy2(t *testing.T) {
	user := &User{
		ID:   1,
		Name: "test",
		Image: &Image{
			ID:  2,
			Url: "test",
		},
	}

	userInfo := &UserInfo{}
	StructCopy(user, userInfo)
	t.Log(userInfo)
}
