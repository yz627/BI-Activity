package home

import (
	"bi-activity/dao"
	"bi-activity/dao/home"
	"bi-activity/models"
	"bi-activity/models/label"
	"bi-activity/response/errors"
	"bi-activity/utils"
	"context"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
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

func (as *ActivityService) GetActivityDetail(ctx context.Context, id uint) (*Activity, error) {
	if id <= 0 {
		return nil, errors.ParameterNotValid
	}
	// 更新浏览量
	_ = as.rr.UpdateActivityViewCount(ctx, id)

	// 获取活动信息
	info, err := as.ar.GetActivityInfoByID(ctx, id)
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
		CreatedAt:                info.CreatedAt.Format(time.DateTime),
		ActivityStatus:           info.ActivityStatus,
	}, nil
}

func (as *ActivityService) SearchActivity(ctx context.Context, params SearchActivityParams) (list []*ActivityCard, err error) {
	if err := as.isValidSearchParams(params); err != nil {
		return nil, err
	}

	daoParams := home.SearchParams{
		ActivityPublisherID: params.ActivityPublisherID,
		ActivityDateEnd:     params.ActivityDateEnd,
		ActivityDateStart:   params.ActivityDateStart,
		ActivityNature:      params.ActivityNature,
		ActivityStatus:      params.ActivityStatus,
		ActivityTypeID:      params.ActivityTypeID,
		Keyword:             params.Keyword,
		Page:                params.Page,
	}

	var activities []*models.Activity
	switch params.ActivityPublisherID {
	case 0: // 全部活动
		activities, err = as.ar.SearchActivity(ctx, daoParams)
	default: // 我的活动
		activities, err = as.ar.SearchMyActivity(ctx, daoParams)
	}

	if err != nil {
		return nil, errors.SearchActivityError
	}

	for _, item := range activities {
		publisherName, err := as.ar.GetPublisherNameByID(ctx, item.ID)
		if err != nil {
			return nil, errors.GetActivityInfoErrorType1
		}

		remainingNumber, err := as.ar.GetActivityRemainingNumberByID(ctx, item.ID)
		if err != nil {
			return nil, errors.GetActivityInfoErrorType4
		}

		list = append(list, &ActivityCard{
			ID:                    item.ID,
			ActivityName:          item.ActivityName,
			ActivityDate:          utils.TransTimeToDate(item.ActivityDate),
			StartTime:             utils.TransTimeToHour(item.StartTime),
			EndTime:               utils.TransTimeToHour(item.EndTime),
			ActivityTypeName:      item.ActivityType.TypeName,
			ActivityTypeImageUrl:  item.ActivityType.Image.URL,
			ActivityPublisherName: publisherName,
			RemainingNumber:       remainingNumber,
		})
	}

	return list, nil
}

// 判断是合法的查询条件
func (as *ActivityService) isValidSearchParams(params SearchActivityParams) error {
	// 活动状态判断
	if !as.isValidActivityStatus(params.ActivityStatus) {
		return errors.SearchActivityParamsErrorType1
	}

	// 活动性质判断
	if !as.isValidActivityNature(params.ActivityNature) {
		return errors.SearchActivityParamsErrorType2
	}

	// 判断时间是否合法
	if !as.isValidActivityDate(params.ActivityDateStart, params.ActivityDateEnd) {
		return errors.SearchActivityParamsErrorType3
	}

	// 活动标签ID合法
	if params.ActivityTypeID <= 0 {
		return errors.ParameterNotValid
	}

	return nil
}

// isValidActivityStatus 判断活动状态是否合法
func (as *ActivityService) isValidActivityStatus(status int) bool {
	switch status {
	case 0:
		return true
	case label.ActivityStatusProceeding, label.ActivityStatusRecruiting, label.ActivityStatusEnded:
		return true
	default:
		return false
	}
}

// isValidActivityNature 判断活动性质是否合法
func (as *ActivityService) isValidActivityNature(nature int) bool {
	switch nature {
	case 0:
		return true
	case label.ActivityNatureStudent, label.ActivityNatureCollege:
		return true
	default:
		return false
	}
}

// isValidActivityDate 判断活动日期是否合法
// 前端限制返回日期不为空，默认为当前时间
func (as *ActivityService) isValidActivityDate(start, end string) bool {
	// 1. 是合法的日期格式
	parse1, err := time.Parse(time.DateOnly, start)
	if err != nil {
		return false
	}

	parse2, err := time.Parse(time.DateOnly, end)
	if err != nil {
		return false
	}
	// 2. 活动开始时间在结束时间之前
	return parse1.Equal(parse2) || parse1.Before(parse2)
}
