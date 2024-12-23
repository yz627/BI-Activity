package dao

import (
	"context"
	"github.com/sirupsen/logrus"
	"strconv"
)

// TODO: PopularActivityKey 更改为配置文件

var PopularActivityKey = "popular_activity" // 热门活动的zset的key

type RedisRepo interface {
	UpdateActivityViewCount(ctx context.Context, id uint) error
	// GetPopularActivities 获取热门活动, 获取前20个
	GetPopularActivities(ctx context.Context) ([]string, error)
}

type redisDataCase struct {
	rdb *Redis
	key string
	log *logrus.Logger
}

func NewRedisDataCase(rdb *Redis, key string, logger *logrus.Logger) RedisRepo {
	return &redisDataCase{
		rdb: rdb,
		key: key,
		log: logger,
	}
}

// UpdateActivityViewCount 向redis中更新活动-浏览量的数据
// key 	 为活动ID
// value 为活动信息
// 存储为redis的有序集合： zset
// 当zset中存在该key时，更新该key的value，否则添加该key-value
// zset 中只保存前20 个元素
func (r *redisDataCase) UpdateActivityViewCount(ctx context.Context, id uint) error {
	_, err := r.rdb.rdb.ZIncrBy(ctx, PopularActivityKey, 1, strconv.Itoa(int(id))).Result()

	// 只保留前20个元素
	_, err = r.rdb.rdb.ZRemRangeByRank(ctx, PopularActivityKey, 20, -1).Result()
	if err != nil {
		r.log.Error("redis remove activity data error:", err)
	}

	return err
}

func (r *redisDataCase) GetPopularActivities(ctx context.Context) ([]string, error) {
	result, err := r.rdb.rdb.ZRevRange(ctx, PopularActivityKey, 0, -1).Result()
	if err != nil {
		r.log.Error("redis get popular activities error:", err)
		return nil, err
	}
	return result, nil
}
