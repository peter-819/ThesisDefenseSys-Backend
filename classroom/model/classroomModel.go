package model

import (
	"context"
	"time"
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
	"TDS-backend/common/mongox"
	"TDS-backend/common/errorx"
)
type Classroom struct{
	ID primitive.ObjectID `bson:"_id"`
	Name string `bson:"name"`
	Location string `bson:"location"`
	StartTime time.Time `bson:"start_time"`
	EndTime time.Time `bson:"end_time"`
}

type IClassroomModel interface {
	RemoveClassroom(id string) error
	InsertClassroom(classroom *Classroom)(error)
	FindClassroomByIdHex(id primitive.ObjectID)(*Classroom, error)
	FindClassroomByIdString(id string)(*Classroom, error)
	QueryAllClassrooms()([]Classroom, error)
	QueryAvailableByTime(start time.Time, end time.Time) ([]Classroom, error)
}

type ClassroomModel struct {
	collection *mongo.Collection
}

func NewClassroomModel(d *mongox.Database) IClassroomModel {
	return &ClassroomModel{
		collection: d.Conn.Collection("Classroom"),
	}
}

func (m *ClassroomModel) RemoveClassroom(id string) error{
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()

	hexId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errorx.NewDefaultError("非法教室ID")
	}

	filter := bson.M{ "_id": hexId }
	_, err = m.collection.DeleteOne(ctx, filter)
	if err != nil {
		return errorx.NewDefaultError("删除教室失败: " + err.Error())
	}
	return nil
}

func (m *ClassroomModel) InsertClassroom(classroom *Classroom) (error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()
	classroom.ID = primitive.NewObjectID()
	_, err := m.collection.InsertOne(ctx, classroom)
	if err != nil {
		return errorx.NewDefaultError("创建教室失败:" + err.Error())
	}
	return nil
}

func (m *ClassroomModel) FindClassroomByIdHex(id primitive.ObjectID)(*Classroom, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()

    filter := bson.M{"_id": id}
	res := &Classroom{}
    if err := m.collection.FindOne(ctx, filter).Decode(res); err != nil {
        return res, errorx.NewDefaultError("查询教室失败: " + err.Error())
    }
	return res, nil
}

func (m *ClassroomModel) FindClassroomByIdString(id string)(*Classroom, error) {
	hexId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errorx.NewDefaultError("非法教室ID")
	}
	return m.FindClassroomByIdHex(hexId)
}

func (m *ClassroomModel) QueryAllClassrooms()([]Classroom, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()
	
	filter := bson.M{}
	cur, err := m.collection.Find(ctx, filter)
	if err != nil {
		return []Classroom{}, errorx.NewDefaultError("数据库错误")
	}
	var result []Classroom
	cur.All(ctx, &result)
	return result, nil
}

func (m *ClassroomModel) QueryAvailableByTime(start time.Time, end time.Time) ([]Classroom, error){
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()
	
	filter := bson.M{
		"start_time" : bson.M{
			"$lte":start,
		},
		"end_time": bson.M{
			"$gte":end,
		},
	}
	cur, err := m.collection.Find(ctx, filter)
	if err != nil {
		return nil, errorx.NewDefaultError("数据库错误: " + err.Error())
	}
	var result []Classroom
	cur.All(ctx, &result)
	return result, nil
}