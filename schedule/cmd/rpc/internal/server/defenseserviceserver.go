// Code generated by goctl. DO NOT EDIT!
// Source: defense.proto

package server

import (
	"context"

	"TDS-backend/schedule/cmd/rpc/internal/logic"
	"TDS-backend/schedule/cmd/rpc/internal/svc"
	"TDS-backend/schedule/cmd/rpc/types/defense"
)

type DefenseServiceServer struct {
	svcCtx *svc.ServiceContext
	defense.UnimplementedDefenseServiceServer
}

func NewDefenseServiceServer(svcCtx *svc.ServiceContext) *DefenseServiceServer {
	return &DefenseServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *DefenseServiceServer) QueryAllDefenses(ctx context.Context, in *defense.EmptyRequest) (*defense.QueryAllDefensesReply, error) {
	l := logic.NewQueryAllDefensesLogic(ctx, s.svcCtx)
	return l.QueryAllDefenses(in)
}
