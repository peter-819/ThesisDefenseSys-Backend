package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UserRpc zrpc.RpcClientConf
	Mongodb struct {
		DataSource   string
		DatabaseName string
	}
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
}
