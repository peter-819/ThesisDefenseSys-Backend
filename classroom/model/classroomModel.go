package model

import (
	"TDS-backend/common/errorx"
	"TDS-backend/common/mongox"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Defense struct {
	DefenseId string    `bson:"defense_id" json:"defense_id"`
	StartTime time.Time `bson:"start_time" json:"start_time"`
	EndTime   time.Time `bson:"end_time" json:"end_time"`
}
type Classroom struct {
	Id        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Location  string             `bson:"location" json:"location"`
	StartTime time.Time          `bson:"start_time" json:"start_time"`
	EndTime   time.Time          `bson:"end_time" json:"end_time"`
	Defenses  []Defense          `bson:"defenses" json:"defenses"`
}

type IClassroomModel interface {
	RemoveClassroom(id string) error
	AddClassroom(classroom *Classroom) error
	FindClassroomByIdHex(id primitive.ObjectID) (*Classroom, error)
	FindClassroomByIdString(id string) (*Classroom, error)
	QueryAllClassrooms() ([]Classroom, error)
	ModifyClassroom(id string, classroom *Classroom) error
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
func (m *ClassroomModel) ModifyClassroom(id string, classroom *Classroom) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	hexId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errorx.NewDefaultError("修改教室ID错误: " + err.Error())
	}
	if !classroom.Id.IsZero() && id != classroom.Id.Hex() {
		return errorx.NewDefaultError("修改教室失败: 不可以修改ID")
	}
	filter := bson.M{"_id": hexId}
	update := bson.M{
		"$set": classroom,
	}
	update_res := m.collection.FindOneAndUpdate(ctx, filter, update)
	if update_res.Err() != nil {
		return errorx.NewDefaultError("教室修改失败: " + update_res.Err().Error())
	}
	return nil
}
func (m *ClassroomModel) RemoveClassroom(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	hexId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errorx.NewDefaultError("非法教室ID")
	}

	filter := bson.M{"_id": hexId}
	_, err = m.collection.DeleteOne(ctx, filter)
	if err != nil {
		return errorx.NewDefaultError("删除教室失败: " + err.Error())
	}
	return nil
}

func (m *ClassroomModel) AddClassroom(classroom *Classroom) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	classroom.Id = primitive.NewObjectID()
	_, err := m.collection.InsertOne(ctx, classroom)
	if err != nil {
		return errorx.NewDefaultError("创建教室失败:" + err.Error())
	}
	return nil
}

func (m *ClassroomModel) FindClassroomByIdHex(id primitive.ObjectID) (*Classroom, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	res := &Classroom{}
	if err := m.collection.FindOne(ctx, filter).Decode(res); err != nil {
		return res, errorx.NewDefaultError("查询教室失败: " + err.Error())
	}
	return res, nil
}

func (m *ClassroomModel) FindClassroomByIdString(id string) (*Classroom, error) {
	hexId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errorx.NewDefaultError("非法教室ID")
	}
	return m.FindClassroomByIdHex(hexId)
}

func (m *ClassroomModel) QueryAllClassrooms() ([]Classroom, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
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

func (m *ClassroomModel) QueryAvailableByTime(start time.Time, end time.Time) ([]Classroom, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{
		"start_time": bson.M{
			"$lte": start,
		},
		"end_time": bson.M{
			"$gte": end,
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
