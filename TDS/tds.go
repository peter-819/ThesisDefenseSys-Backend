package main

import (
	"flag"
	"fmt"

	"TDS-backend/TDS/internal/config"
	"TDS-backend/TDS/internal/handler"
	"TDS-backend/TDS/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"TDS-backend/TDS/internal/data"
)

var configFile = flag.String("f", "etc/tds-api.yaml", "the config file")

func main() {
	go serverdata.GetServerData()
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
