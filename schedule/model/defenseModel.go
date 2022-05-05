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

type Member struct {
	TeacherID   string `bson:"teacher_id" json:"teacher_id"`
	TeacherName string `bson:"teacher_name" json:"teacher_name"`
}
type Student struct {
	StudentID   string `bson:"student_id" json:"student_id"`
	StudentName string `bson:"student_name" json:"student_name"`
}
type Defense struct {
	Id                primitive.ObjectID `bson:"_id" json:"id"`
	StartTime         time.Time          `bson:"start_time" json:"start_time"`
	EndTime           time.Time          `bson:"end_time" json:"end_time"`
	Classroom         string             `bson:"classroom" json:"classroom"`
	ClassroomFullName string             `bson:"classroom_full_name" json:"classroom_full_name"`
	Committee         []Member           `bson:"committee" json:"committee"`
	Students          []Student          `bson:"students" json:"students"`
}

type IDefenseModel interface {
	InsertDefense(defense *Defense) (string, error)
	RemoveDefense(id string) error
	ModifyDefense(id string, newDefense *Defense) error
	GetAllDefense() ([]Defense, error)
	GetTeacherDefense(id string) ([]Defense, error)
	GetStudentDefense(id string) ([]Defense, error)
	QueryDefense(id string) (*Defense, error)

	IsClassroomInUse(id string) (bool, error)
}

type DefenseModel struct {
	collection *mongo.Collection
}

func NewDefenseModel(d *mongox.Database) IDefenseModel {
	return &DefenseModel{
		collection: d.Conn.Collection("Defense"),
	}
}

func (m *DefenseModel) InsertDefense(defense *Defense) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	defense.Id = primitive.NewObjectID()
	res, err := m.collection.InsertOne(ctx, defense)
	if err != nil {
		return "", errorx.NewDefaultError("创建答辩失败: " + err.Error())
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (m *DefenseModel) RemoveDefense(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	hexId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errorx.NewDefaultError("删除答辩失败: 答辩 ID 格式错误")
	}

	filter := bson.M{"_id": hexId}
	_, err = m.collection.DeleteOne(ctx, filter)
	if err != nil {
		return errorx.NewDefaultError("删除答辩失败: " + err.Error())
	}
	return nil
}

func (m *DefenseModel) ModifyDefense(id string, newDefense *Defense) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	hexId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errorx.NewDefaultError("修改答辩失败: 答辩 ID 格式错误")
	}
	newDefense.Id = hexId
	filter := bson.M{"_id": hexId}
	update := bson.M{"$set": newDefense}
	_, err = m.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return errorx.NewDefaultError("修改答辩失败: " + err.Error())
	}
	return nil
}

func (m *DefenseModel) GetAllDefense() ([]Defense, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{}
	cur, err := m.collection.Find(ctx, filter)
	if err != nil {
		return []Defense{}, errorx.NewDefaultError("数据库错误")
	}
	var result []Defense
	cur.All(ctx, &result)
	return result, nil
}

func (m *DefenseModel) GetStudentDefense(id string) ([]Defense, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{
		"student_id": id,
	}
	cur, err := m.collection.Find(ctx, filter)
	if err != nil {
		return []Defense{}, errorx.NewDefaultError("数据库错误")
	}
	var result []Defense
	cur.All(ctx, &result)
	return result, nil
}

func (m *DefenseModel) GetTeacherDefense(id string) ([]Defense, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{
		"committee": bson.M{
			"$elemMatch": bson.M{
				"teacher_id": id,
			},
		},
	}
	cur, err := m.collection.Find(ctx, filter)
	if err != nil {
		return []Defense{}, errorx.NewDefaultError("数据库错误")
	}
	var result []Defense
	cur.All(ctx, &result)
	return result, nil
}

func (m *DefenseModel) IsClassroomInUse(id string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{
		"classroom": id,
	}
	count, err := m.collection.CountDocuments(ctx, filter)
	if err != nil {
		return false, errorx.NewDefaultError("数据库错误: " + err.Error())
	}
	if count >= 1 {
		return true, nil
	}
	return false, nil
}

func (m *DefenseModel) QueryDefense(id string) (*Defense, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	hexId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errorx.NewDefaultError("非法答辩ID")
	}

	filter := bson.M{"_id": hexId}
	res := &Defense{}
	if err := m.collection.FindOne(ctx, filter).Decode(res); err != nil {
		return res, errorx.NewDefaultError("查询答辩失败: " + err.Error())
	}
	return res, nil
}
