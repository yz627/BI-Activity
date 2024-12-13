package dao

import (
	"bi-activity/configs"
	"bi-activity/global"
	"bi-activity/models"
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Data struct {
	Db *gorm.DB
}

type Redis struct {
	Rdb redis.Cmdable
	RDB *redis.Client
}

func NewDataDao(c *configs.Database, logger *logrus.Logger) *Data {
	db, err := gorm.Open(mysql.Open(c.DSN()), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用外键约束
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 设置表名的映射采取单数
		},
	})

	if err != nil {
		logger.Fatalf("mysql connect error: %v", err)
	}

	// 执行数据库迁移
	err = db.AutoMigrate(
		&models.Student{},            // 自动创建 student 表
		&models.Image{},              // 自动创建 image 表
		&models.College{},            // 自动创建 college 表
		&models.Activity{},           // 自动创建 activity 表
		&models.Participant{},        // 自动创建 participant 表
		&models.StudentActivityAudit{}, // 自动创建 student_activity_audit 表
	)
	if err != nil {
		logger.Fatalf("database migration failed: %v", err)
	}

	global.Db = db

	return &Data{Db: db}
}

func NewRedisDao(c *configs.Redis, logger *logrus.Logger) *Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Addr,                                     // redis地址
		ReadTimeout:  time.Second * time.Duration(c.ReadTimeout),               // 读取超时时间
		WriteTimeout: time.Second * time.Duration(c.WriteTimeout),              // 写入超时时间
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
		Rdb: rdb,
		RDB: rdb,
	}
}
