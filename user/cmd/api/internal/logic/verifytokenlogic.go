package logic

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"TDS-backend/common/errorx"
	"TDS-backend/user/cmd/api/internal/svc"
	"TDS-backend/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyTokenLogic {
	return &VerifyTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type Role int64

const (
	Anyone    = 1 << iota
	Student   = 1<<iota | Anyone
	Teacher   = 1<<iota | Student
	Secretary = 1<<iota | Teacher
	Admin     = 1<<iota | Secretary
)

func getRole(role string) Role {
	switch role {
	case "Admin":
		return Admin
	case "Secretary":
		return Secretary
	case "Teacher":
		return Teacher
	case "Student":
		return Student
	default:
		panic("invalid role")
	}
}
func matchUrl(url string, prefixs []string) bool {
	for _, pre := range prefixs {
		if strings.HasPrefix(url, pre) {
			return true
		}
	}
	return false
}
func check(role Role, cRole Role) bool {
	fmt.Println("checking: ", role, cRole)
	return (cRole | role) > 0
}
func (l *VerifyTokenLogic) VerifyToken(req *types.VerifyTokenReq, r *http.Request) (resp *types.VerifyTokenResp, err error) {
	// todo: add your logic here and delete this line
	authorization := r.Header.Get("Authorization")
	realRequestPath := r.Header.Get("X-Original-Uri")
	fmt.Println("authorization: ", authorization)
	fmt.Println("realRequestPath: ", realRequestPath)
	roleString := l.ctx.Value("role").(string)
	fmt.Println("role string: ", roleString)
	id := l.ctx.Value("id").(string)

	role := getRole(roleString)
	fmt.Println("role enum: ", role)

	if matchUrl(realRequestPath, l.svcCtx.Config.AdminUrls) && check(role, Admin) {
		fmt.Println("Admin pass")
		return &types.VerifyTokenResp{
			UserId:   id,
			UserRole: roleString,
			Ok:       true,
		}, nil
	}
	if matchUrl(realRequestPath, l.svcCtx.Config.SecretaryUrls) && check(role, Secretary) {
		fmt.Println("Secretary pass")
		return &types.VerifyTokenResp{
			UserId:   id,
			UserRole: roleString,
			Ok:       true,
		}, nil
	}
	if matchUrl(realRequestPath, l.svcCtx.Config.TeacherUrls) && check(role, Teacher) {
		fmt.Println("Teacher pass")
		return &types.VerifyTokenResp{
			UserId:   id,
			UserRole: roleString,
			Ok:       true,
		}, nil
	}
	if matchUrl(realRequestPath, l.svcCtx.Config.StudentUrls) && check(role, Student) {
		fmt.Println("Student pass")
		return &types.VerifyTokenResp{
			UserId:   id,
			UserRole: roleString,
			Ok:       true,
		}, nil
	}
	if matchUrl(realRequestPath, l.svcCtx.Config.AnyoneUrls) && check(role, Anyone) {
		fmt.Println("Anyone pass")
		return &types.VerifyTokenResp{
			UserId:   id,
			UserRole: roleString,
			Ok:       true,
		}, nil
	}
	urls := []string{}
	urls = append(urls, l.svcCtx.Config.AdminUrls...)
	urls = append(urls, l.svcCtx.Config.SecretaryUrls...)
	urls = append(urls, l.svcCtx.Config.TeacherUrls...)
	urls = append(urls, l.svcCtx.Config.StudentUrls...)
	urls = append(urls, l.svcCtx.Config.AnyoneUrls...)
	if matchUrl(realRequestPath, urls) {
		fmt.Println("role error")
		return nil, errorx.NewCodeError(401, "没有访问权限")
	} else {
		fmt.Println("url error")
		return nil, errorx.NewCodeError(404, "无效地址")
	}
}
