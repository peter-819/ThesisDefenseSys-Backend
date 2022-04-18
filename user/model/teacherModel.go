package model 

import (
	"context"
	"time"
	
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"TDS-backend/common/mongox"
	"TDS-backend/common/errorx"
)

type Teacher struct {
	ID string `bson:"id"`
	Name string `bson:"name"`
	IsSecretary string `bson:"is_secretary"`
}

type ITeacherModel interface {
	GetTeacherList() ([]Teacher, error)
	GetTeacher(id string)(Teacher, error)
	SetTeacherInfo(id string, info Teacher)error
}

type TeacherModel struct {
	collection *mongo.Collection
}

func NewTeacherModel(d *mongox.Database) ITeacherModel {
	return &TeacherModel{
		collection: d.Conn.Collection("TeacherCollection"),
	}
}

func (m *TeacherModel) GetTeacherList() ([]Teacher, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	filter := bson.M{}
	
	cur, find_err := m.collection.Find(ctx, filter)
	if find_err != nil {
		return []Teacher{}, errorx.NewDefaultError("数据库错误")
	} else {
		var results []Teacher
		cur.All(ctx, &results)
		return results, nil
	}
}

func (m *TeacherModel) GetTeacher(id string)(Teacher, error){
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var result Teacher
	find_err := m.collection.FindOne(ctx,bson.D{
		{"id", id},
	}).Decode(&result)
	
	if find_err != nil {
		if find_err == mongo.ErrNoDocuments {
			return result, errorx.NewDefaultError("用户不存在")
		} else {
			return result, errorx.NewDefaultError("数据库错误")
		}
	}

	return result, nil
}

func (m *TeacherModel) SetTeacherInfo(id string, info Teacher) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()

	filter := bson.M{"id": id}
	
	update := bson.M{
		"$set":bson.M{},
	}
	if info.Name != "" {
		update["$set"].(bson.M)["name"] = info.Name
	}
	if info.IsSecretary != "" {
		if info.IsSecretary == "yes" {
			update["$set"].(bson.M)["is_secretary"] = info.IsSecretary
		} else if info.IsSecretary == "no" {
			update["$set"].(bson.M)["is_secretary"] = info.IsSecretary
		}
	}
	// if info.ID != "" {
		// update["$set"].(bson.M)["id"] = info.ID
	// }
	update_res := m.collection.FindOneAndUpdate(ctx,filter,update)
	if update_res.Err() != nil {
		// fmt.Println(update_res.Err())
		return errorx.NewDefaultError("教师信息修改失败: " + update_res.Err().Error())
	}
	return nil
}