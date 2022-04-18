package utils

import (
	"TDS-backend/student/cmd/api/internal/types"
	"TDS-backend/student/model"
)

func StudentJtoB(in *types.Student) (out *model.Student, err error) {
	out = &model.Student{
		Id:            in.Id,
		Name:          in.Name,
		PaperKeywords: in.PaperKeywords,
		PaperTitle:    in.PaperTitle,
		GroupId:       in.GroupId,
		Mentor:        in.Mentor,
	}
	return out, nil
}

func StudentBtoJ(in *model.Student) (out *types.Student) {
	out = &types.Student{
		Id:            in.Id,
		Name:          in.Name,
		PaperKeywords: in.PaperKeywords,
		PaperTitle:    in.PaperTitle,
		GroupId:       in.GroupId,
		Mentor:        in.Mentor,
	}
	return out
}
