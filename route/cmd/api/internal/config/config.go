package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}	
	Mongodb struct {
		DataSource string
		DatabaseName string
	}	
}
