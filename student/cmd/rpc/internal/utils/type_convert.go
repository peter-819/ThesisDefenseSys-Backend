package utils

import (
	"TDS-backend/common/errorx"
	"TDS-backend/student/cmd/rpc/types/student"
	"TDS-backend/student/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func StudentJtoB(in *student.Student) (out *model.Student, err error) {
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

func StudentBtoJ(in *model.Student) (out *student.Student) {
	out = &student.Student{
		Id:            in.Id,
		Name:          in.Name,
		PaperKeywords: in.PaperKeywords,
		PaperTitle:    in.PaperTitle,
		GroupId:       in.GroupId,
		Mentor:        in.Mentor,
	}
	return out
}

func GroupJtoB(in *student.Group) (out *model.Group, err error) {
	var hexId primitive.ObjectID
	if in.Id != "" {
		hexId, err = primitive.ObjectIDFromHex(in.Id)
		if err != nil {
			return nil, errorx.NewDefaultError("教室 ID 格式错误: " + err.Error())
		}
	} else {
		hexId = primitive.NewObjectID()
	}
	out = &model.Group{
		Id:        hexId,
		Name:      in.Name,
		DefenseId: in.DefenseId,
		Members:   in.Members,
	}
	return out, nil
}

func GroupBtoJ(in *model.Group) (out *student.Group) {
	out = &student.Group{
		Id:        in.Id.Hex(),
		Name:      in.Name,
		DefenseId: in.DefenseId,
		Members:   in.Members,
	}
	return out
}
