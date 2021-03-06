package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	ClassroomRpc zrpc.RpcClientConf
	TeacherRpc   zrpc.RpcClientConf
	StudentRpc   zrpc.RpcClientConf
	Mongodb      struct {
		DataSource   string
		DatabaseName string
	}
}
