package dao

import (
	"context"
	"github.com/sirupsen/logrus"
)

type RedisRepo interface {
	UpdateActivityViewCount(ctx context.Context, key string, value interface{}) error
	GetPopularActivities(ctx context.Context) ([]string, error)
}

type redisDataCase struct {
	rdb *Redis
	log *logrus.Logger
}

func NewRedisDataCase(rdb *Redis, logger *logrus.Logger) RedisRepo {
	return &redisDataCase{
		rdb: rdb,
		log: logger,
	}
}

// UpdateActivityViewCount 向redis中更新活动-浏览量的数据
// key 	 为活动ID
// value 为活动信息
// 存储为redis的有序集合： zset
func (r *redisDataCase) UpdateActivityViewCount(ctx context.Context, key string, value interface{}) error {
	// TODO: 存储为redis的有序集合： zset
	panic("implement me")
}

func (r *redisDataCase) GetPopularActivities(ctx context.Context) ([]string, error) {
	//TODO implement me
	panic("implement me")
}
