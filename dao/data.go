package dao

import (
	"bi-activity/configs"
	_ "bi-activity/models"
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Data struct {
	db *gorm.DB
}

type Redis struct {
	rdb redis.Cmdable
	RDB *redis.Client
}

// NewDateDao MySQL数据库连接实例
// *Data 数据库连接实例
// func() 函数返回值，返回一个函数，用于释放资源
func NewDateDao(c *configs.Database, logger *logrus.Logger) (*Data, func()) {
	db, err := gorm.Open(mysql.Open(c.DSN()), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用外键约束
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 设置表名的映射采取单数
		},
	})

	if err != nil {
		logger.Fatalf("mysql connect error: %v", err)
	}

	return &Data{db: db}, func() {
		logger.Info("closing the data resources")

		sqlDB, err := db.DB()
		if err != nil {
			logger.Errorf("close db error: %v", err)
		}

		if err := sqlDB.Close(); err != nil {
			logger.Errorf("close db error: %v", err)
		}
	}
}

func (d *Data) DB() *gorm.DB {
	return d.db
}

// NewRedisDao Redis连接实例
// *Redis redis连接实例
// func() 函数返回值，返回一个函数，用于释放资源
func NewRedisDao(c *configs.Redis, logger *logrus.Logger) (*Redis, func()) {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Addr,                                      // redis地址
		ReadTimeout:  time.Second * time.Duration(c.ReadTimeout),  // 读取超时时间
		WriteTimeout: time.Second * time.Duration(c.WriteTimeout), // 写入超时时间
		DialTimeout:  time.Second * time.Duration(c.DialTimeout),  // 连接超时时间
		PoolSize:     c.PoolSize,                                  // 连接池大小
	})
	// 测试连接
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelFunc()
	err := rdb.Ping(timeout).Err()

	// redis 连接失败，退出程序
	if err != nil {
		logger.Fatalf("redis connect error: %v", err)
	}
	return &Redis{
			rdb: rdb,
		}, func() {
			logger.Info("closing the redis resources")

			if err := rdb.Close(); err != nil {
				logger.Errorf("close redis error: %v", err)
			}
		}
}
