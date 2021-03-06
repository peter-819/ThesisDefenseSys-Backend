// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"TDS-backend/teacher/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/query/:id",
				Handler: QueryTeacherHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/modify",
				Handler: ModifyTeacherHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/query/all",
				Handler: QueryAllTeachersHandler(serverCtx),
			},
		},
		rest.WithPrefix("/teacher"),
	)
}
