package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Mongodb struct {
		DataSource string
		DatabaseName string
	}
}
