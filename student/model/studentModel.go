package model

import (
	"TDS-backend/common/errorx"
	"TDS-backend/common/mongox"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Student struct {
	Id            string   `bson:"id"`
	Name          string   `bson:"name"`
	PaperKeywords []string `bson:"paper_keywords"`
	PaperTitle    string   `bson:"paper_title"`
	GroupId       string   `bson:"group_id"`
	Mentor        string   `bson:"mentor"`
}

type IStudentModel interface {
	QueryStudent(id string) (*Student, error)
	RemoveStudent(id string) error
	QueryAllStudents() ([]Student, error)
	QueryNongroupedStudents() ([]Student, error)
	QueryStudentsBatch(ids []string) ([]Student, error)
	ModifyStudent(id string, student *Student) error
	AddStudent(student *Student) error
}

type StudentModel struct {
	collection *mongo.Collection
}

func NewStudentModel(d *mongox.Database) IStudentModel {
	return &StudentModel{
		collection: d.Conn.Collection("StudentCollection"),
	}
}

func (m *StudentModel) QueryStudent(id string) (*Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result := &Student{}
	filter := bson.D{{"id", id}}
	find_err := m.collection.FindOne(ctx, filter).Decode(result)

	if find_err != nil {
		return nil, errorx.NewDefaultError("数据库错误: " + find_err.Error())
	}
	return result, nil
}

func (m *StudentModel) QueryStudentsBatch(ids []string) ([]Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{
		"id": bson.M{
			"$in": ids,
		},
	}
	cur, err := m.collection.Find(ctx, filter)
	if err != nil {
		return nil, errorx.NewDefaultError("数据库错误: " + err.Error())
	}
	var result []Student
	err = cur.All(ctx, &result)
	if err != nil {
		return nil, errorx.NewDefaultError("数据库错误: " + err.Error())
	}
	return result, nil
}

func (m *StudentModel) RemoveStudent(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{"id": id}
	_, err := m.collection.DeleteOne(ctx, filter)
	if err != nil {
		return errorx.NewDefaultError("删除学生失败: " + err.Error())
	}
	return nil
}

func (m *StudentModel) QueryAllStudents() ([]Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.D{}
	cur, err := m.collection.Find(ctx, filter)
	if err != nil {
		return nil, errorx.NewDefaultError("数据库错误: " + err.Error())
	}
	var result []Student
	err = cur.All(ctx, &result)
	if err != nil {
		return nil, errorx.NewDefaultError("数据库错误: " + err.Error())
	}
	return result, nil
}

func (m *StudentModel) QueryNongroupedStudents() ([]Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{
		"$or": []bson.M{
			{
				"group_id": "",
			},
			{
				"group_id": bson.M{
					"$exists": false,
				},
			},
		},
	}
	cur, err := m.collection.Find(ctx, filter)
	if err != nil {
		return nil, errorx.NewDefaultError("数据库错误: " + err.Error())
	}
	var result []Student
	err = cur.All(ctx, &result)
	if err != nil {
		return nil, errorx.NewDefaultError("数据库错误: " + err.Error())
	}
	return result, nil
}

func (m *StudentModel) ModifyStudent(id string, student *Student) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{"id": id}
	update := bson.M{
		"$set": student,
	}
	update_res := m.collection.FindOneAndUpdate(ctx, filter, update)
	if update_res.Err() != nil {
		return errorx.NewDefaultError("学生修改失败: " + update_res.Err().Error())
	}
	return nil
}

func (m *StudentModel) AddStudent(student *Student) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := m.collection.InsertOne(ctx, student)
	if err != nil {
		return errorx.NewDefaultError("创建学生失败: " + err.Error())
	}
	return nil
}
