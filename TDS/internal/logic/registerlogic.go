package logic

import (
	"context"
	"fmt"
    "time"

	"TDS-backend/TDS/internal/svc"
	"TDS-backend/TDS/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req types.RegisterReq) (resp *types.RegisterReply, err error) {
	collection := l.svcCtx.Database.Collection("UserInfoCollection")
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	record := &types.DbUserInfo{
		Username: req.Username,
		Password: req.Password,
		Token: req.Username + "token",
		Role: req.Role,
	}
	res, err := collection.InsertOne(ctx, record)
	id := res.InsertedID
	fmt.Println(id)

    return &types.RegisterReply{
		ErrorCode: 0,
        Message: "Register Success",
    }, nil
}
