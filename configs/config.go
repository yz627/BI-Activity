package configs

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	configPath = ".\\..\\configs"
)

// Config 全局配置信息
type Config struct {
	Database   *Database   `yaml:"Database"`
	Redis      *Redis      `yaml:"Redis"`
	Server     *Server     `yaml:"Server"`
	UserStatus *UserStatus `yaml:"UserStatus"`
}

func InitConfig(path ...string) *Config {
	if len(path) > 0 {
		configPath = path[0]
	}

	viper.AddConfigPath(configPath)
	viper.SetConfigName("configs")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatalf("config file load failed: %s", err)
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		logrus.Fatalf("config file unmarshal failed: %s", err)
	}
	return &config
}

// Database 数据库配置映射
type Database struct {
	Network  string `yaml:"network"`  // 网络类型
	Addr     string `yaml:"addr"`     // 数据库地址
	User     string `yaml:"user"`     // 数据库用户名
	Password string `yaml:"password"` // 数据库密码
	DB       string `yaml:"db"`       // 数据库名
}

func (d *Database) DSN() string {
	return fmt.Sprintf(
		"%s:%s@%s(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		d.User, d.Password, d.Network, d.Addr, d.DB,
	)
}

// Redis redis 配置映射
type Redis struct {
	Network      string `yaml:"network"`                                    // 网络类型
	Addr         string `yaml:"addr"`                                       // 地址
	Password     string `yaml:"password"`                                   // 密码
	ReadTimeout  int    `yaml:"read_timeout" mapstructure:"read_timeout"`   // 读超时
	WriteTimeout int    `yaml:"write_timeout" mapstructure:"write_timeout"` // 写超时
	DialTimeout  int    `yaml:"dial_timeout" mapstructure:"dial_timeout"`   // 连接超时时间
	PoolSize     int    `yaml:"pool_size" mapstructure:"pool_size"`         // 连接池大小
}

// Server 服务配置映射
type Server struct {
	Name string `yaml:"name"`
	Port string `yaml:"port"`
}

// UserStatus 登录状态映射
type UserStatus struct {
	ExpirationTime int64  `yaml:"expiration_time" mapstructure:"expiration_time"` // 登录过期时间
	LoginFlag      string `yaml:"login_flag" mapstructure:"login_flag"`           // 登录标识
}
