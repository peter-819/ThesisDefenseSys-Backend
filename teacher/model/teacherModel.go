package model

import (
	"TDS-backend/common/errorx"
	"TDS-backend/common/mongox"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Schedule struct {
	ScheduleId string    `bson:"schedule_id"`
	Name       string    `bson:"name"`
	StartTime  time.Time `bson:"start_time"`
	EndTime    time.Time `bson:"end_time"`
}
type Teacher struct {
	Id                string     `bson:"id,optional"`
	Name              string     `bson:"name,optional"`
	IsSecretary       string     `bson:"is_secretary"`
	MaxDefensePerWeek int        `bson:"max_defense_per_week"`
	PreferKeywords    []string   `bson:"prefer_keywords"`
	Schedules         []Schedule `bson:"schedules"`
}

type ITeacherModel interface {
	// GetTeacherSchedule(id string) ([]Schedule, error)
	// ResetTeacherSchedule(id string, newSchedules []Schedule) error

	QueryTeacher(id string) (*Teacher, error)
	ModifyTeacher(id string, teacher *Teacher) error
	AddTeacher(teacher *Teacher) error
	QueryAvailableTeachers(start time.Time, end time.Time, keywords []string, excluded []string) ([]Teacher, error)
}

type TeacherModel struct {
	collection *mongo.Collection
}

func NewTeacherModel(d *mongox.Database) ITeacherModel {
	return &TeacherModel{
		collection: d.Conn.Collection("TeacherCollection"),
	}
}

// func (m *TeacherModel) GetTeacherSchedule(id string) ([]Schedule, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
// 	defer cancel()

// 	var result struct {
// 		Schedules []Schedule `bson:"schedules"`
// 	}
// 	filter := bson.D{{"id", id}}
// 	find_err := m.collection.FindOne(ctx, filter).Decode(&result)

// 	if find_err != nil {
// 		return []Schedule{}, errorx.NewDefaultError("数据库错误")
// 	}
// 	return result.Schedules, nil
// }

// func (m *TeacherModel) ResetTeacherSchedule(id string, newSchedules []Schedule) error{
// 	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
// 	defer cancel()

// 	filter := bson.M{"id": id}
// 	update := bson.M{
// 		"$set":bson.M{"schedules" : newSchedules},
// 	}
// 	update_res := m.collection.FindOneAndUpdate(ctx,filter,update)
// 	if update_res.Err() != nil {
// 		// fmt.Println(update_res.Err())
// 		return errorx.NewDefaultError("教师日程重设失败: " + update_res.Err().Error())
// 	}
// 	return nil
// }

func (m *TeacherModel) QueryTeacher(id string) (*Teacher, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result := &Teacher{}
	filter := bson.D{{"id", id}}
	find_err := m.collection.FindOne(ctx, filter).Decode(result)

	if find_err != nil {
		return nil, errorx.NewDefaultError("数据库错误: " + find_err.Error())
	}
	return result, nil
}

func (m *TeacherModel) ModifyTeacher(id string, teacher *Teacher) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{"id": id}
	update := bson.M{
		"$set": teacher,
	}
	update_res := m.collection.FindOneAndUpdate(ctx, filter, update)
	if update_res.Err() != nil {
		return errorx.NewDefaultError("教师修改失败: " + update_res.Err().Error())
	}
	return nil
}

func (m *TeacherModel) AddTeacher(teacher *Teacher) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := m.collection.InsertOne(ctx, teacher)
	if err != nil {
		return errorx.NewDefaultError("创建教师失败: " + err.Error())
	}
	return nil
}

func (m *TeacherModel) QueryAvailableTeachers(start time.Time, end time.Time, keywords []string, excluded []string) ([]Teacher, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{
		"prefer_keywords": bson.M{
			"$elemMatch": bson.M{
				"$in": keywords,
			},
		},
		"id": bson.M{
			"$not": bson.M{
				"$in": excluded,
			},
		},
	}
	if !start.IsZero() && !end.IsZero() {
		filter["schedules"] = bson.M{
			"$not": bson.M{
				"$elemMatch": bson.M{
					"$or": []bson.M{
						bson.M{
							"start_time": bson.M{
								"$gte": start,
								"$lt":  end,
							},
						},
						bson.M{
							"end_time": bson.M{
								"$gt":  start,
								"$lte": end,
							},
						},
					},
				},
			},
		}
	}
	cur, err := m.collection.Find(ctx, filter)
	if err != nil {
		return nil, errorx.NewDefaultError("数据库错误: " + err.Error())
	}
	var result []Teacher
	err = cur.All(ctx, &result)
	if err != nil {
		return nil, errorx.NewDefaultError("数据库错误: " + err.Error())
	}
	return result, nil
}
