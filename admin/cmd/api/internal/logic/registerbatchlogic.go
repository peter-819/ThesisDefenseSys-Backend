package logic

import (
	"context"

	"TDS-backend/admin/cmd/api/internal/svc"
	"TDS-backend/admin/cmd/api/internal/types"
	"TDS-backend/user/cmd/rpc/types/user"
	"TDS-backend/excel/cmd/rpc/types/excel"
	"TDS-backend/common/errorx"
	

	"mime/multipart"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"sort"
	"fmt"
)
// var COL_NAMES = [...]string{"NameRow", "IdRow", "PasswordRow", "RoleRow"} 
type RegisterInfo struct {
	RowMin int32
	RowMax int32

	Cols map[string]int32
	DefaultRole string
	DefaultPassword string
}
type RegisterbatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	
	Files []*multipart.FileHeader
	Info RegisterInfo
}

func NewRegisterbatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterbatchLogic {
	return RegisterbatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
type Pair struct {
	Key string
	Value int32
}
func (l *RegisterbatchLogic) Registerbatch() (resp *types.RegisterBatchReply, err error) {
	resp = &types.RegisterBatchReply{}
	if l.ctx.Value("role") != "Admin" {
		return resp, errorx.NewDefaultError("没有权限")
	}
	file, err := l.Files[0].Open()
	defer file.Close()
	if err != nil {
		return resp, errorx.NewDefaultError("文件打开错误")
	}
	// data := make([]byte, FileMaxLength)
	data, err := io.ReadAll(file) //TODO
	if err != nil {
		return resp, errorx.NewDefaultError("文件读取错误")
	}

	stream, err := l.svcCtx.ExcelRpc.ParseExcel(l.ctx)
	cols := []int32{}
	pairs := []Pair{}
	for k,v := range l.Info.Cols {
		cols = append(cols, v)
		pairs = append(pairs, Pair{Key:k,Value:v,})
	}
	sort.SliceStable(pairs, func(i,j int) bool {
		return pairs[i].Value < pairs[j].Value
	})

	stream.Send(&excel.ExcelRequest{
		Info: &excel.ExcelInfo{
			RowMin: l.Info.RowMin,
			RowMax: l.Info.RowMax,
			Cols: cols,
		},
	})
	excel.ChunkerSrv(data).ChunkerSend(stream)
	res, err := stream.CloseAndRecv()
	if err != nil {
		return resp, errorx.NewDefaultError("rpc 接收错误: " + err.Error())
	}
	
	userInfo := &user.RegisterBatchRequest{}
	for _, rowInfo := range res.Rows{
		req := &user.RegisterRequest{}
		row := rowInfo.Elements
		if v, ok := l.Info.Cols["IdCol"]; ok && row[v] != "" { 
			req.Id = row[v] 
		} else {
			return resp, errorx.NewDefaultError("ID列号错误或有空值!")
		} 
		if v, ok := l.Info.Cols["NameCol"]; ok && row[v] != "" { 
			req.Name = row[v]
		} else {
			return resp, errorx.NewDefaultError("名字列号错误或有空值！")
		}

		if v, ok := l.Info.Cols["RoleCol"]; ok && row[v] != "" { 
			req.Role = row[v]
		} else if l.Info.DefaultRole != "" { 
			req.Role = l.Info.DefaultRole 
		} else {
			return resp, errorx.NewDefaultError("角色错误或有空值！")
		}

		if v, ok := l.Info.Cols["PasswordCol"]; ok && row[v] != "" { 
			req.Password = row[v]
		} else if l.Info.DefaultPassword != "" { 
			req.Password = l.Info.DefaultPassword
		} else {
			return resp, errorx.NewDefaultError("密码错误或有空值！")
		}
		userInfo.List = append(userInfo.List, req)
	}
	fmt.Println(userInfo)
	_, err = l.svcCtx.UserRpc.RegisterBatch(l.ctx, userInfo)
	if err != nil {
		return resp, err
	}
	return resp,nil 
}
