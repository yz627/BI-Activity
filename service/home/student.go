package home

import (
	"bi-activity/dao/home"
	"bi-activity/response/errors"
	"context"
	"github.com/sirupsen/logrus"
)

type StudentService struct {
	sr  home.StudentRepo
	log *logrus.Logger
}

func NewStudentService(sr home.StudentRepo, logger *logrus.Logger) *StudentService {
	return &StudentService{
		sr:  sr,
		log: logger,
	}
}

func (ss *StudentService) StudentInfo(ctx context.Context, id uint) (*StuInfo, error) {
	if id <= 0 {
		return nil, errors.ParameterNotValid
	}

	resp, err := ss.sr.GetStudentInfoByID(ctx, id)
	if err != nil {
		return nil, errors.GetStudentInfoByIDError
	}

	return &StuInfo{
		CollegeName: resp.College.CollegeName,
		Email:       resp.StudentEmail,
		ID:          resp.StudentID,
		Name:        resp.StudentName,
		Phone:       resp.StudentPhone,
	}, nil
}
