package configs

import "time"

type Config struct {
	Database *Database `yaml:"Database"`
	Redis    *Redis    `yaml:"Redis"`
	Server   *Server   `yaml:"Server"`
}

type Database struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	User string `yaml:"user"`
	Pass string `yaml:"password"`
}

type Redis struct {
	Network      string         `yaml:"network"`
	Addr         string         `yaml:"addr"`
	Password     string         `yaml:"password"`
	ReadTimeout  *time.Duration `yaml:"read_timeout"`
	WriteTimeout *time.Duration `yaml:"write_timeout"`
}

type Server struct {
	Name string `yaml:"name"`
	Port string `yaml:"port"`
}
