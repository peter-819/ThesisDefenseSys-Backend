// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"TDS-backend/TDS/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/route/menu",
				Handler: routemenuHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/route/header",
				Handler: routeheaderHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: loginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: registerHandler(serverCtx),
			},
		},
	)
}
