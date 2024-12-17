package dao

import (
	"context"
	"github.com/sirupsen/logrus"
)

type RedisRepo interface {
	UpdateActivityViewCount(ctx context.Context, data interface{}, count int) error
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

func (r *redisDataCase) UpdateActivityViewCount(ctx context.Context, data interface{}, count int) error {
	//TODO implement me
	panic("implement me")
}

func (r *redisDataCase) GetPopularActivities(ctx context.Context) ([]string, error) {
	//TODO implement me
	panic("implement me")
}
