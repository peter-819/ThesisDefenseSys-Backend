package utils

import (
	"TDS-backend/common/errorx"
	"TDS-backend/common/timex"
	"TDS-backend/schedule/cmd/api/internal/types"
	"TDS-backend/schedule/model"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DefenseJtoB(in *types.Defense) (out *model.Defense, err error) {
	members := []model.Member{}
	for _, m := range in.Committee {
		members = append(members, model.Member{
			TeacherID:   m.TeacherID,
			TeacherName: m.TeacherName,
		})
	}
	var id primitive.ObjectID
	fmt.Println("id: " + in.Id)
	if in.Id != "" {
		id, err = primitive.ObjectIDFromHex(in.Id)
		if err != nil {
			return nil, errorx.NewDefaultError("答辩 ID 格式错误: " + err.Error())
		}
	} else {
		id = primitive.NewObjectID()
	}

	start_time, err := timex.FStringToTime(in.StartTime)
	if err != nil {
		return nil, errorx.NewDefaultError("答辩开始时间格式错误: " + err.Error())
	}
	end_time, err := timex.FStringToTime(in.StartTime)
	if err != nil {
		return nil, errorx.NewDefaultError("答辩开始时间格式错误: " + err.Error())
	}

	out = &model.Defense{
		Id:                id,
		StartTime:         start_time,
		EndTime:           end_time,
		Classroom:         in.Classroom,
		ClassroomFullName: in.ClassroomFullName,
		Committee:         members,
		GroupId:           in.GroupId,
	}
	return out, nil
}

func DefenseBtoJ(in *model.Defense) (out *types.Defense) {
	members := []types.Member{}
	for _, m := range in.Committee {
		members = append(members, types.Member{
			TeacherID:   m.TeacherID,
			TeacherName: m.TeacherName,
		})
	}
	out = &types.Defense{
		Id:                in.Id.Hex(),
		StartTime:         timex.TimeToString(in.StartTime),
		EndTime:           timex.TimeToString(in.EndTime),
		Classroom:         in.Classroom,
		ClassroomFullName: in.ClassroomFullName,
		Committee:         members,
		GroupId:           in.GroupId,
	}
	return out
}
