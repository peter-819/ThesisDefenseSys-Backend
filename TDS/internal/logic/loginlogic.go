package logic

import (
	"context"
    "time"
	"fmt"

	"TDS-backend/TDS/internal/svc"
	"TDS-backend/TDS/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	jwt "github.com/dgrijalva/jwt-go"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req types.LoginReq) (resp *types.LoginReply, err error) {
	collection := l.svcCtx.Database.Collection("UserInfoCollection")
	fmt.Println(req)
	
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	var result types.DbUserInfo

	find_err := collection.FindOne(ctx,bson.D{
		{"username", req.Username},
	}).Decode(&result)
	if find_err == mongo.ErrNoDocuments {
		return &types.LoginReply{
			ErrorCode: 1,
			Message: "Invalid Username",
		}, nil
	} else if find_err != nil{
		return &types.LoginReply{
			ErrorCode: 404,
			Message: "Internal Error",
		}, nil
	} else {
		if result.Password != req.Password {
			return &types.LoginReply{
				ErrorCode: 2,
				Message: "Wrong Password",
			}, nil
		} else{
			now := time.Now().Unix()
			token, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, result.Username )
			if err != nil {
				fmt.Println("failed to gen token")
			}
			fmt.Println("gen token .... : ", token)
			return &types.LoginReply{
				ErrorCode: 0,
				Message: "Login Success",
				Token: token,
			},nil
		}
	}
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds int64, username string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["username"] = username
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
  }