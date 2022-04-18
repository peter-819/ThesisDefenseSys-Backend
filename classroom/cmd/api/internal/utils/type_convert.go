package utils

import (
	"TDS-backend/common/errorx"
	"time"
	"TDS-backend/classroom/cmd/api/internal/types"
	"TDS-backend/classroom/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ClassroomJtoB(in *types.Classroom) (out *model.Classroom, err error) {
	var id primitive.ObjectID
	if in.ID != "" {
		id, err = primitive.ObjectIDFromHex(in.ID)
		if err != nil {
			return nil, errorx.NewDefaultError("教室 ID 格式错误: " + err.Error())
		}
	} else {
		id = primitive.NewObjectID()
	}

	start_time, err := time.Parse("2006-01-02T15:04:05", in.StartTime)
	if err != nil {
		return nil, errorx.NewDefaultError("教室开始时间格式错误: " + err.Error())
	}
	end_time, err := time.Parse("2006-01-02T15:04:05", in.EndTime)
	if err != nil {
		return nil, errorx.NewDefaultError("教室结束时间格式错误: " + err.Error())
	}	

	out = &model.Classroom{
		ID:id,
		Name: in.Name,
		StartTime: start_time,
		EndTime: end_time,
		Location: in.Location,
	}
	return out, nil
}

func ClassroomBtoJ(in *model.Classroom) (out *types.Classroom) {
	out = &types.Classroom{
		ID: in.ID.Hex(),
		Name: in.Name,
		StartTime: in.StartTime.String(),
		EndTime: in.EndTime.String(),
		Location: in.Location,
	}
	return out
}