package homeSvc

import (
	"bi-activity/dao/homeDao"
	"bi-activity/response/errors"
	"context"
	"github.com/sirupsen/logrus"
)

type BiDataService struct {
	ar  homeDao.ActivityRepo
	sr  homeDao.StudentRepo
	cr  homeDao.CollegeRepo
	log *logrus.Logger
}

func NewBiDataService(ar homeDao.ActivityRepo, sr homeDao.StudentRepo, cr homeDao.CollegeRepo, log *logrus.Logger) *BiDataService {
	return &BiDataService{
		ar:  ar,
		sr:  sr,
		cr:  cr,
		log: log,
	}
}

func (bs *BiDataService) BiData(ctx context.Context) (*BiData, error) {
	activityTotal, err := bs.ar.GetActivityTotal(ctx)
	if err != nil {
		return nil, errors.GetActivityTotalError
	}

	collegeTotal, err := bs.cr.GetCollegeTotal(ctx)
	if err != nil {
		return nil, errors.GetCollegeTotalError
	}

	studentTotal, err := bs.sr.GetStudentTotal(ctx)
	if err != nil {
		return nil, errors.GetStudentTotalError
	}

	return &BiData{
		ActivityTotal: activityTotal,
		CollegeTotal:  collegeTotal,
		StudentTotal:  studentTotal,
	}, nil
}

func (bs *BiDataService) BiDataLeaderboard(ctx context.Context) ([]*BiDataLeaderboard, error) {
	// 获取学院-人数的映射关系
	res, err := bs.sr.GetCollegeStudentCount(ctx)
	if err != nil {
		return nil, errors.GetCollegeStudentCountError
	}

	var list []*BiDataLeaderboard
	for _, item := range res {
		list = append(list, &BiDataLeaderboard{
			CollegeName:  item.CollegeName,
			StudentTotal: item.Count,
		})
	}

	return list, nil
}
