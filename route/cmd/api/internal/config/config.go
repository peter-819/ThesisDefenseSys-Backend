package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Mongodb struct {
		DataSource   string
		DatabaseName string
	}
}
