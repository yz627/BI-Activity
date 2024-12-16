package service

import (
	"bi-activity/dao"
	"bi-activity/response/errors"
	"context"
	"github.com/sirupsen/logrus"
	"time"
)

type Activity struct {
	// 活动信息
	ID                      uint
	ActivityNature          int       // 活动类别 1-学生活动 2-学院活动
	ActivityStatus          int       // 活动状态 1-未开始 2-进行中 3-已结束
	ActivityName            string    // 活动名称
	ActivityTypeName        uint      // 活动类型，需要查询活动类型，通过ID转化
	ActivityAddress         string    // 活动地址
	ActivityIntroduction    string    // 活动内容
	ActivityDate            time.Time // 活动时间
	StartTime               time.Time // 活动开始时间
	EndTime                 time.Time // 活动结束时间
	RecruitmentNumber       uint      // 活动人数
	RecruitmentRestriction  int       // 活动限制 1-无限制 2-学院内
	RecruitmentRequirements string    // 活动要求
	RecruitmentDeadline     time.Time // 活动截止时间
	ContactName             string    // 活动联系人姓名
	ContactDetails          string    // 活动联系人联系方式

	// 活动图片信息
	ImageID uint
	Url     string

	// 活动发起人信息
	PublisherID   uint
	PublisherName string
}

type ActivityType struct {
	ID       uint
	TypeName string
	ImageID  uint
	Url      string
}

type ActivityService struct {
	ar  dao.ActivityRepo
	log *logrus.Logger
}

func NewActivityService(ar dao.ActivityRepo, log *logrus.Logger) *ActivityService {
	return &ActivityService{
		ar:  ar,
		log: log,
	}
}

func (as *ActivityService) ActivityAllTypes(ctx context.Context) (list []*ActivityType, err error) {
	// 获取所有活动类型
	resp, err := as.ar.GetActivityAllTypes(ctx)
	if err != nil {
		return nil, errors.GetActivityTypeError
	}

	for _, item := range resp {
		list = append(list, &ActivityType{
			ID:       item.ID,
			TypeName: item.TypeName,
			ImageID:  item.ImageID,
			Url:      item.Image.Url,
		})
	}
	return list, nil
}

// PopularActivity 获取热门活动
func (as *ActivityService) PopularActivity(ctx context.Context) (list []*Activity, err error) {
	// TODO: implement me
	panic("implement me")
}

func (as *ActivityService) ActivityInformation(ctx context.Context, id uint) (*Activity, error) {
	if id <= 0 {
		return nil, errors.ParameterNotValid
	}

	// TODO: implement me
	panic("implement me")
}
