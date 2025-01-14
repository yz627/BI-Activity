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
	// TODO：ir未使用
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
		return nil, errors.TypeGetAllActivityTypeError
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
	// 从redis中获取热门活动id
	result, err := as.rr.GetPopularActivities(ctx)
	if err != nil {
		return nil, errors.ActivityPopularActivityError
	}

	// 将result转换为uint列表
	activityID := make([]uint, 0, len(result))
	for _, item := range result {
		id, _ := strconv.Atoi(item)
		activityID = append(activityID, uint(id))
	}

	// 获取活动信息
	// 这里dao接口同时查询了三个表：图片表、活动类型表、活动表
	activityList, err := as.ar.GetActivityListByID(ctx, activityID)
	if err != nil {
		return nil, errors.ActivityPopularActivityError
	}

	// 数据模型的转换DO -> DTO
	// TODO: 更合适在控制层完成
	for _, item := range activityList {
		list = append(list, &ActivityCard{
			ID:                   item.ID,
			ActivityName:         item.ActivityName,
			ActivityDate:         parse.TransTimeToDate(item.ActivityDate),
			StartTime:            parse.TransTimeToHour(item.StartTime),
			EndTime:              parse.TransTimeToHour(item.EndTime),
			ActivityTypeName:     item.ActivityType.TypeName,
			ActivityTypeImageUrl: item.ActivityType.Image.URL,
		})
	}

	// 获取活动发布人信息
	for i, item := range list {
		name, err := as.ar.GetPublisherNameByID(ctx, item.ID)
		if err != nil {
			return nil, errors.ActivityPublisherInfoError
		}
		list[i].ActivityPublisherName = name
	}

	// 为了满足页面展示，需要按照activityID的顺序重新排序
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

// GetActivityDetail 获取活动详情
// TODO: 石山代码，需要大量重构
func (as *ActivityService) GetActivityDetail(ctx context.Context, aID, sID uint) (*Activity, error) {
	// 活动ID需要合法
	if aID <= 0 {
		return nil, errors.ActivityIdNotValid
	}

	// 获取活动信息
	info, err := as.ar.GetActivityInfoByID(ctx, aID)
	if err != nil {
		return nil, errors.ActivityDetailsInfoError
	}
	// 如果活动不存在，返回错误
	if info.ID == 0 {
		return nil, errors.ActivityNotFoundError
	}

	// 更新浏览量
	err = as.rr.UpdateActivityViewCount(ctx, aID)
	if err != nil {
		as.log.Errorf("update activity page view err: %v", err)
	}

	// 获取活动发布人信息
	publisherName, err := as.ar.GetPublisherNameByID(ctx, info.ID)
	if err != nil {
		return nil, errors.ActivityPublisherInfoError
	}

	// 获取活动报名人数
	enrollNumber, err := as.ar.GetActivityEnrollNumberByID(ctx, info.ID)
	if err != nil {
		return nil, errors.ActivityDetailsInfoError
	}

	// 如果当前为学生登陆状态，获取该学生的报名状态
	status := 0
	if sID > 0 {
		status, _ = as.ar.GetParticipateStatus(ctx, sID, aID)
	}

	// 判断当前学生是否为活动发布者
	isPublisher := 0
	if info.ActivityPublisherID == sID {
		isPublisher = 1
	}

	// 判断是否满足活动限制
	isFull := 1
	if sID > 0 && info.RegistrationRestrictions == label.RecruitmentRestrictionCollege {
		// 获取学生的院系ID
		collegeID, _ := as.ar.GetStudentCollegeID(ctx, sID)
		// 1. 发布者学院，比较发布者ID和当前登录用户的院系ID是否一致
		if info.ActivityNature == label.ActivityNatureCollege {
			if collegeID != info.ActivityPublisherID {
				isFull = 0
			}
		}
		// 2. 发布者学生, 比较当前登录用户的院系ID和活动发布者的院系ID是否一致
		if info.ActivityNature == label.ActivityNatureStudent {
			// 获取学生的院系ID
			publisherCollegeID, _ := as.ar.GetStudentCollegeID(ctx, info.ActivityPublisherID)
			if publisherCollegeID != collegeID {
				isFull = 0
			}
		}
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
		IsPublisher:              isPublisher,
		CreatedAt:                info.CreatedAt.Format(time.DateTime),
		ActivityStatus:           info.ActivityStatus,
		ParticipateStatus:        status,
		IsCompliance:             isFull,
	}, nil
}

func (as *ActivityService) SearchActivity(ctx context.Context, params SearchActivityParams) (list []*ActivityCard, count int64, err error) {
	// 参数校验
	if err := as.isValidSearchParams(params); err != nil {
		return nil, -1, err
	}

	// 转换为dao层需要的参数结构
	daoParams := homeDao.SearchParams{
		// ActivityPublisherID为0表示不在登录状态，查询全部活动，否则需要查询我的活动
		ActivityPublisherID: params.ActivityPublisherID,
		ActivityDateEnd:     params.ActivityDateEnd,
		ActivityDateStart:   params.ActivityDateStart,
		ActivityNature:      params.ActivityNature,
		ActivityStatus:      params.ActivityStatus,
		ActivityTypeID:      params.ActivityTypeID,
		Keyword:             params.Keyword,
		Page:                params.Page,
	}

	// 获取活动列表
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
			return nil, -1, errors.ActivityDetailsInfoError
		}

		remainingNumber, err := as.ar.GetActivityRemainingNumberByID(ctx, item.ID)
		if err != nil {
			return nil, -1, errors.ActivityDetailsInfoError
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
		return errors.SearchParamsNotValid
	}

	// 活动性质判断
	if !as.isValidActivityNature(params.ActivityNature) {
		return errors.SearchParamsNotValid
	}

	// 判断时间是否合法
	if !as.isValidActivityDate(params.ActivityDateStart, params.ActivityDateEnd) {
		return errors.SearchParamsNotValid
	}

	// 活动标签ID合法
	if params.ActivityTypeID < 0 {
		return errors.SearchParamsNotValid
	}

	// 页数判断
	if params.Page < 1 {
		return errors.SearchParamsNotValid
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
func (as *ActivityService) isValidActivityDate(start, end string) bool {
	// 都不为空
	// 1. 是合法的日期格式--time.DateOnly格式
	// 2. 活动开始时间在结束时间之前
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

		return parse1.Equal(parse2) || parse1.Before(parse2)
	}

	// 开始时间为空
	// 1. 是合法的日期格式
	if start == "" && end != "" {
		_, err := time.Parse(time.DateOnly, end)
		if err != nil {
			return false
		}
	}

	// 结束时间为空
	// 1. 是合法的日期格式
	if end == "" && start != "" {
		_, err := time.Parse(time.DateOnly, start)
		if err != nil {
			return false
		}
	}

	// 都为空--合法
	return true
}

func (as *ActivityService) ParticipateActivity(ctx context.Context, stuID, activityID uint) error {
	if stuID <= 0 || activityID <= 0 {
		return errors.ParticipateParamsNotValid
	}

	// 判断是否复合报名条件
	// 1. 活动处于招募状态
	// 2. 活动的招募条件为全体学生
	// 3. 活动的招募条件为学院学生，且学生所在学院为报名活动的学院
	// 4. 报名者是活动发布者, 无需报名

	// 判断是否已经报名
	status, err := as.ar.GetParticipateStatus(ctx, stuID, activityID)
	if err != nil {
		return errors.ParticipateStatusError
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

func (as *ActivityService) EditActivityType(ctx context.Context, id int, name string) error {
	if id <= 0 {
		return errors.TypeEditTypeIdError
	}

	if name == "" {
		return errors.TypeEditTypeNameError
	}

	return as.tr.UpdateActivityTypeByID(ctx, id, name)
}

func (as *ActivityService) DeleteActivityType(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.TypeEditTypeIdError
	}

	return as.tr.DeleteActivityTypeByID(ctx, id)
}

func (as *ActivityService) AddActivityType(ctx context.Context, imageId int, typeName string) (*models.ActivityType, error) {
	if imageId <= 0 {
		return nil, errors.TypeEditTypeIdError
	}

	return as.tr.AddActivityType(ctx, imageId, typeName)
}
