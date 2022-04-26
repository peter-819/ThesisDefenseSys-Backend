package logic

import (
	"context"

	"TDS-backend/common/errorx"
	"TDS-backend/student/cmd/rpc/internal/svc"
	"TDS-backend/student/cmd/rpc/types/student"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryStudentsContentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryStudentsContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryStudentsContentLogic {
	return &QueryStudentsContentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryStudentsContentLogic) QueryStudentsContent(in *student.QueryStudentsContentRequest) (*student.QueryGroupContentResponse, error) {
	// todo: add your logic here and delete this line
	resp := &student.QueryGroupContentResponse{}
	keywordsMap := map[string]byte{}
	mentorsMap := map[string]byte{}

	for _, m := range in.Ids {
		stu, err := l.svcCtx.StudentModel.QueryStudent(m)
		if err != nil {
			return nil, errorx.NewDefaultError("学生不存在: " + err.Error())
		}
		for _, keyword := range stu.PaperKeywords {
			keywordsMap[keyword] = 0
		}
		mentorsMap[stu.Mentor] = 0
	}
	for key, _ := range keywordsMap {
		if key == "" {
			continue
		}
		resp.Keywords = append(resp.Keywords, key)
	}
	for key, _ := range mentorsMap {
		if key == "" {
			continue
		}
		resp.Mentors = append(resp.Mentors, key)
	}
	return resp, nil
}
