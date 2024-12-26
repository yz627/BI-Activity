package homeSvc

import (
	"bi-activity/dao"
	"bi-activity/dao/homeDao"
	"bi-activity/models"
	"bi-activity/models/label"
	"bi-activity/response/errors"
	"bi-activity/utils/parse"
	"context"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type ActivityService struct {
	ar  homeDao.ActivityRepo
	ir  homeDao.ImageRepo
	tr  homeDao.ActivityTypeRepo
	rr  dao.RedisRepo
	log *logrus.Logger
}

func NewActivityService(ar homeDao.ActivityRepo, ir homeDao.ImageRepo, tr homeDao.ActivityTypeRepo, rr dao.RedisRepo, log *logrus.Logger) *ActivityService {
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
// TODO: 耗时较长，需要优化
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
	activityList, err := as.ar.GetActivityListByID(ctx, activityID)
	if err != nil {
		return nil, errors.GetActivityError
	}

	for _, item := range activityList {
		list = append(list, &ActivityCard{
			ID:                   item.ID,
			ActivityName:         item.ActivityName,
			ActivityDate:         parse.TransTimeToDate(item.ActivityDate),
			StartTime:            parse.TransTimeToHour(item.StartTime),
			EndTime:              parse.TransTimeToHour(item.EndTime),
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

	// 按照activityID的顺序重新排序
	activityMap := make(map[uint]*ActivityCard, len(activityID))
	for _, item := range list {
		activityMap[item.ID] = item
	}

	list = make([]*ActivityCard, len(activityID))
	for i, id := range activityID {
		list[i] = activityMap[id]
	}

	return list, nil
}

func (as *ActivityService) GetActivityDetail(ctx context.Context, aID, sID uint) (*Activity, error) {
	// TODO：添加活动参加显示，如果已经成功参加活动，需要在活动详情中显示

	if aID <= 0 {
		return nil, errors.ParameterNotValid
	}

	// 获取活动信息
	info, err := as.ar.GetActivityInfoByID(ctx, aID)
	if err != nil {
		return nil, errors.GetActivityInfoErrorType2
	}
	// 如果活动不存在，返回错误
	if info.ID == 0 {
		return nil, errors.GetActivityInfoError
	}

	// 更新浏览量
	_ = as.rr.UpdateActivityViewCount(ctx, aID)

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

	// 获取报名状态
	status := 0
	if sID > 0 {
		status, _ = as.ar.GetParticipateStatus(ctx, sID, aID)
	}

	return &Activity{
		ID:                       info.ID,
		ActivityAddress:          info.ActivityAddress,
		ContactName:              info.ContactName,
		ContactDetails:           info.ContactDetails,
		ActivityTypeName:         info.ActivityType.TypeName,
		ActivityTypeImageUrl:     info.ActivityType.Image.URL,
		ActivityDate:             parse.TransTimeToDate(info.ActivityDate),
		StartTime:                parse.TransTimeToHour(info.StartTime),
		EndTime:                  parse.TransTimeToHour(info.EndTime),
		RecruitmentNumber:        info.RecruitmentNumber,
		RecruitedNumber:          enrollNumber,
		RegistrationRestrictions: label.RecruitmentRestriction[info.RegistrationRestrictions],
		RegistrationRequirement:  info.RegistrationRequirement,
		RegistrationDeadline:     parse.TransTimeToTime(info.RegistrationDeadline),
		ActivityIntroduction:     info.ActivityIntroduction,
		ActivityContent:          info.ActivityContent,
		ActivityName:             info.ActivityName,
		ActivityImageUrl:         info.ActivityImage.URL,
		PublisherName:            publisherName,
		CreatedAt:                info.CreatedAt.Format(time.DateTime),
		ActivityStatus:           info.ActivityStatus,
		ParticipateStatus:        status,
	}, nil
}

func (as *ActivityService) SearchActivity(ctx context.Context, params SearchActivityParams) (list []*ActivityCard, count int64, err error) {
	if err := as.isValidSearchParams(params); err != nil {
		return nil, -1, err
	}

	as.log.Debugf("SearchActivity params: %+v", params)

	daoParams := homeDao.SearchParams{
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
		activities, count, err = as.ar.SearchActivity(ctx, daoParams)
	default: // 我的活动
		activities, count, err = as.ar.SearchMyActivity(ctx, daoParams)
	}
	if err != nil {
		return nil, -1, errors.SearchActivityError
	}

	for _, item := range activities {
		publisherName, err := as.ar.GetPublisherNameByID(ctx, item.ID)
		if err != nil {
			return nil, -1, errors.GetActivityInfoErrorType1
		}

		remainingNumber, err := as.ar.GetActivityRemainingNumberByID(ctx, item.ID)
		if err != nil {
			return nil, -1, errors.GetActivityInfoErrorType4
		}

		list = append(list, &ActivityCard{
			ID:                    item.ID,
			ActivityName:          item.ActivityName,
			ActivityDate:          parse.TransTimeToDate(item.ActivityDate),
			StartTime:             parse.TransTimeToHour(item.StartTime),
			EndTime:               parse.TransTimeToHour(item.EndTime),
			ActivityTypeName:      item.ActivityType.TypeName,
			ActivityTypeImageUrl:  item.ActivityType.Image.URL,
			ActivityPublisherName: publisherName,
			RemainingNumber:       remainingNumber,
		})
	}

	return list, count, nil
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
	if params.ActivityTypeID < 0 {
		return errors.ParameterNotValid
	}

	// 页数判断
	if params.Page < 1 {
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
	if start != "" && end != "" {
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

	if start == "" && end != "" {
		_, err := time.Parse(time.DateOnly, end)
		if err != nil {
			return false
		}
	}

	if end == "" && start != "" {
		_, err := time.Parse(time.DateOnly, start)
		if err != nil {
			return false
		}
	}
	return true
}

func (as *ActivityService) ParticipateActivity(ctx context.Context, stuID, activityID uint) error {
	if stuID <= 0 || activityID <= 0 {
		return errors.ParameterNotValid
	}

	// 判断是否复合报名条件
	// 1. 活动处于招募状态
	// 2. 活动的招募条件为全体学生
	// 3. 活动的招募条件为学院学生，且学生所在学院为报名活动的学院
	// TODO

	// 判断是否已经报名
	status, err := as.ar.GetParticipateStatus(ctx, stuID, activityID)
	if err != nil {
		return errors.GetParticipateStatusError
	}

	switch status {
	case label.ParticipateStatusPassed:
		return errors.ParticipateActivityErrorType1
	case label.ParticipateStatusPending:
		return errors.ParticipateActivityErrorType2
	}

	// 进行报名
	err = as.ar.AddParticipate(ctx, stuID, activityID)
	if err != nil {
		return errors.ParticipateActivityErrorType3
	}
	return nil
}
