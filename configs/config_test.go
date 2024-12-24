package configs

import "testing"

func TestInitConfig(t *testing.T) {
	conf := InitConfig()
	t.Log(conf.Database)
	t.Log(conf.Redis)
	t.Log(conf.Redis.Password)
	t.Log(conf.Server)
	t.Log(conf.UserStatus)
	t.Log(conf.Server.ServerAddress())
}
