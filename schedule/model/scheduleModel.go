package model

import (
	"context"
	"time"
	
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"TDS-backend/common/mongox"
	"TDS-backend/common/errorx"
)

type TSchedule struct {
	Name string `bson:"name"`
	StartTime time.Time `bson:"start_time"`
	EndTime time.Time `bson:"end_time"`
	ID string `bson:"id"`
}

type DSchedule struct {
	Name string `bson:"name"`
	StartTime time.Time `bson:"start_time"`
	EndTime time.Time `bson:"end_time"`
	Classroom string `bson:"classroom"`
}

type IScheduleModel interface {
	GetTeacherSchedule(id string) ([]TSchedule, error)
	GetTeacherDefense(id string) ([]DSchedule, error)
	GetStudentDefense(id string) ([]DSchedule, error)

	ResetTeacherSchedule(id string, newSchedules []TSchedule) error 
}

type ScheduleModel struct {
	teacherCollection *mongo.Collection
	defenseCollection *mongo.Collection
}

func NewScheduleModel(d *mongox.Database) IScheduleModel {
	return &ScheduleModel{
		teacherCollection: d.Conn.Collection("TeacherCollection"),
		defenseCollection: d.Conn.Collection("DefenseSchedule"),
	}
}

func (m *ScheduleModel) GetTeacherSchedule(id string) ([]TSchedule, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()

	var result struct {
		Schedules []TSchedule `bson:"schedules"`
	}
	filter := bson.D{{"id", id}}
	find_err := m.teacherCollection.FindOne(ctx, filter).Decode(&result)

	if find_err != nil {
		return []TSchedule{}, errorx.NewDefaultError("数据库错误") 
	}
	return result.Schedules, nil
}

func (m *ScheduleModel) GetStudentDefense(id string) ([]DSchedule, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()
	
	filter := bson.M{
		"student_id": id,
	}
	cur, err := m.defenseCollection.Find(ctx, filter)
	if err != nil {
		return []DSchedule{}, errorx.NewDefaultError("数据库错误")
	}
	var result []DSchedule
	cur.All(ctx, &result)
	return result, nil
}

func (m *ScheduleModel) GetTeacherDefense(id string) ([]DSchedule, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()

	filter := bson.M{
		"committee":bson.M{
			"$elemMatch":bson.M{
				"teacher_id": id,
			},
		},
	}
	cur, err := m.defenseCollection.Find(ctx, filter)
	if err != nil {
		return []DSchedule{}, errorx.NewDefaultError("数据库错误")
	}
	var result []DSchedule
	cur.All(ctx, &result)
	return result, nil
}

func (m *ScheduleModel) ResetTeacherSchedule(id string, newSchedules []TSchedule) error{
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()

	filter := bson.M{"id": id}
	update := bson.M{
		"$set":bson.M{"schedules" : newSchedules},
	}
	update_res := m.teacherCollection.FindOneAndUpdate(ctx,filter,update)
	if update_res.Err() != nil {
		// fmt.Println(update_res.Err())
		return errorx.NewDefaultError("教师日程重设失败: " + update_res.Err().Error())
	}
	return nil
}