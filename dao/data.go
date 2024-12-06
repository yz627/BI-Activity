package dao

import (
	"bi-activity/configs"
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type Data struct {
	db *gorm.DB
}

type Redis struct {
	rdb redis.Cmdable
}

func NewDateDao(c *configs.Database, logger *logrus.Logger) *Data {
	db, err := gorm.Open(mysql.Open(c.DSN()), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用外键约束
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 设置表名的映射采取单数
		},
	})

	if err != nil {
		logger.Fatalf("mysql connect error: %v", err)
	}

	return &Data{db: db}
}

func NewRedisDao(c *configs.Redis, logger *logrus.Logger) *Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Addr,                                     // redis地址
		ReadTimeout:  time.Duration(c.ReadTimeout),               // 读取超时时间
		WriteTimeout: time.Duration(c.WriteTimeout),              // 写入超时时间
		DialTimeout:  time.Second * time.Duration(c.DialTimeout), // 连接超时时间
		PoolSize:     c.PoolSize,                                 // 连接池大小
		Password:     c.Password,                                 // 密码
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
	}
}
