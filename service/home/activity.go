package home

import (
	"bi-activity/dao"
	"bi-activity/dao/home"
	"bi-activity/models/label"
	"bi-activity/response/errors"
	"bi-activity/utils"
	"context"
	"github.com/sirupsen/logrus"
	"strconv"
)

type ActivityService struct {
	ar  home.ActivityRepo
	ir  home.ImageRepo
	tr  home.ActivityTypeRepo
	rr  dao.RedisRepo
	log *logrus.Logger
}

func NewActivityService(ar home.ActivityRepo, ir home.ImageRepo, tr home.ActivityTypeRepo, rr dao.RedisRepo, log *logrus.Logger) *ActivityService {
	return &ActivityService{
		ir:  ir,
		tr:  tr,
		ar:  ar,
		rr:  rr,
		log: log,
	}
}

// ActivityAllTypes 获取所有活动类型
func (as *ActivityService) ActivityAllTypes(ctx context.Context) (list []*ActivityType, err error) {
	// 获取所有活动类型
	resp, err := as.tr.GetActivityAllTypes(ctx)
	if err != nil {
		return nil, errors.GetActivityTypeError
	}

	for _, item := range resp {
		list = append(list, &ActivityType{
			ID:       item.ID,
			TypeName: item.TypeName,
			Url:      item.Image.URL,
		})
	}
	return list, nil
}

// PopularActivity 获取热门活动
// 只需要返回部分信息
// 因为热门活动只展示卡片信息，卡片信息只需要展示活动名称、发布人名称、活动时间、活动类型名称、活动类型图片
func (as *ActivityService) PopularActivity(ctx context.Context) (list []*ActivityCard, err error) {
	result, err := as.rr.GetPopularActivities(ctx)
	if err != nil {
		return nil, errors.GetPopularActivityError
	}

	// 将result转换为uint列表
	activityID := make([]uint, 0, len(result))
	for _, item := range result {
		id, _ := strconv.Atoi(item)
		activityID = append(activityID, uint(id))
	}

	// 获取活动信息
	// TODO: 活动发布者在两个表中，必须一个一个获取，无法一次全部查询
	activityList, err := as.ar.GetActivityListByID(ctx, activityID)
	if err != nil {
		return nil, errors.GetActivityError
	}

	for _, item := range activityList {
		list = append(list, &ActivityCard{
			ID:                   item.ID,
			ActivityName:         item.ActivityName,
			ActivityDate:         utils.TransTimeToDate(item.ActivityDate),
			StartTime:            utils.TransTimeToHour(item.StartTime),
			EndTime:              utils.TransTimeToHour(item.EndTime),
			ActivityTypeName:     item.ActivityType.TypeName,
			ActivityTypeImageUrl: item.ActivityType.Image.URL,
			//ActivityPublisherName 发布者名称
			//RemainingNumber       int    // 活动招募剩余人数
		})
	}

	// 获取活动发布人信息
	for i, item := range list {
		name, err := as.ar.GetPublisherNameByID(ctx, item.ID)
		if err != nil {
			return nil, errors.GetActivityInfoErrorType1
		}
		list[i].ActivityPublisherName = name
	}

	return list, nil
}

func (as *ActivityService) GetActivityDetail(ctx context.Context, activityID string) (*Activity, error) {
	if activityID == "" {
		return nil, errors.ParameterNotValid
	}
	id, _ := strconv.Atoi(activityID)

	// 更新浏览量
	_ = as.rr.UpdateActivityViewCount(ctx, uint(id))

	// 获取活动信息
	info, err := as.ar.GetActivityInfoByID(ctx, uint(id))
	if err != nil {
		return nil, errors.GetActivityInfoErrorType2
	}

	// 获取活动发布人信息
	publisherName, err := as.ar.GetPublisherNameByID(ctx, info.ID)
	if err != nil {
		return nil, errors.GetActivityInfoErrorType1
	}

	// 获取活动报名人数
	enrollNumber, err := as.ar.GetActivityEnrollNumberByID(ctx, info.ID)
	if err != nil {
		return nil, errors.GetActivityInfoErrorType3
	}

	return &Activity{
		ID:                       info.ID,
		ActivityAddress:          info.ActivityAddress,
		ContactName:              info.ContactName,
		ContactDetails:           info.ContactDetails,
		ActivityTypeName:         info.ActivityType.TypeName,
		ActivityTypeImageUrl:     info.ActivityType.Image.URL,
		ActivityDate:             utils.TransTimeToDate(info.ActivityDate),
		StartTime:                utils.TransTimeToHour(info.StartTime),
		EndTime:                  utils.TransTimeToHour(info.EndTime),
		RecruitmentNumber:        info.RecruitmentNumber,
		RecruitedNumber:          enrollNumber,
		RegistrationRestrictions: label.RecruitmentRestriction[info.RegistrationRestrictions],
		RegistrationRequirement:  info.RegistrationRequirement,
		RegistrationDeadline:     utils.TransTimeToTime(info.RegistrationDeadline),
		ActivityIntroduction:     info.ActivityIntroduction,
		ActivityContent:          info.ActivityContent,
		ActivityName:             info.ActivityName,
		ActivityImageUrl:         info.ActivityImage.URL,
		PublisherName:            publisherName,
		CreatedAt:                info.CreatedAt.Format("2006-01-02 15:04:05"),
		ActivityStatus:           info.ActivityStatus,
	}, nil
}

func (as *ActivityService) SearchActivity(ctx context.Context, keyword string) (list []*ActivityCard, err error) {
	// TODO: 搜索活动
	panic("implement me")
}
