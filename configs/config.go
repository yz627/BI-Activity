package configs

import "time"

type Config struct{}

type Database struct {
	Host string
	Port string
	User string
	Pass string
}

type Redis struct {
	Network      string
	Addr         string
	Password     string
	ReadTimeout  *time.Duration
	WriteTimeout *time.Duration
}

type Server struct {
}
