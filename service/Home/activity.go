package Home

import (
	"bi-activity/dao"
	"bi-activity/response/errors"
	"context"
	"github.com/sirupsen/logrus"
	"strconv"
)

type ActivityService struct {
	ar  dao.ActivityRepo
	ir  dao.ImageRepo
	tr  dao.ActivityTypeRepo
	rr  dao.RedisRepo
	log *logrus.Logger
}

func NewActivityService(ar dao.ActivityRepo, ir dao.ImageRepo, tr dao.ActivityTypeRepo, rr dao.RedisRepo, log *logrus.Logger) *ActivityService {
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

	// 获取图片url
	listID := make([]uint, 0, len(resp))
	for _, item := range resp {
		listID = append(listID, item.ImageID)
	}
	listUrl, err := as.ir.GetImageUrlsByID(ctx, listID)
	if err != nil {
		return nil, errors.GetImageError
	}

	for i, item := range resp {
		list = append(list, &ActivityType{
			ID:       item.ID,
			TypeName: item.TypeName,
			Url:      listUrl[i],
		})
	}
	return list, nil
}

// PopularActivity 获取热门活动
// 只需要返回部分信息
// 因为热门活动只展示卡片信息，卡片信息只需要展示活动名称和图片即可
// TODO: 优化返回信息
func (as *ActivityService) PopularActivity(ctx context.Context) (list []*Activity, err error) {
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

	// 获取图片url
	imageID := make([]uint, 0, len(activityList))
	for _, item := range activityList {
		imageID = append(imageID, item.ActivityImageID)
	}
	_, err = as.ir.GetImageUrlsByID(ctx, imageID)
	if err != nil {
		return nil, errors.GetImageError
	}

	// 获取活动
	return nil, nil
}
