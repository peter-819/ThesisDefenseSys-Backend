package model

import (
	"context"
	"time"

	"TDS-backend/common/errorx"
	"TDS-backend/common/mongox"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID       string `bson:"id" json:"id"`
	Name     string `bson:"name" json:"name"`
	Password string `bson:"password" json:"password"`
	Role     string `bson:"role" json:"role"`
}

type IUserModel interface {
	GetUser(id string) (*User, error)
	GetTeachers() ([]User, error)
	AddUser(id, name, role, password string) error
	ModifyUser(id string, newUser *User) error
}

type UserModel struct {
	collection *mongo.Collection
}

func NewUserModel(d *mongox.Database) IUserModel {
	return &UserModel{
		collection: d.Conn.Collection("User"),
	}
}

func (m *UserModel) ModifyUser(id string, newUser *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{"id": id}
	update := bson.M{
		"$set": newUser,
	}
	update_res := m.collection.FindOneAndUpdate(ctx, filter, update)
	if update_res.Err() != nil {
		return errorx.NewDefaultError("教师用户失败: " + update_res.Err().Error())
	}
	return nil
}

func (m *UserModel) GetUser(id string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var result User
	find_err := m.collection.FindOne(ctx, bson.D{
		{"id", id},
	}).Decode(&result)

	if find_err != nil {
		if find_err == mongo.ErrNoDocuments {
			return &result, errorx.NewDefaultError("用户不存在")
		} else {
			return &result, errorx.NewDefaultError("数据库错误")
		}
	}

	return &result, nil
}
func (m *UserModel) GetTeachers() ([]User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{
		"$or": []bson.M{
			bson.M{"role": "Teacher"},
			bson.M{"role": "Secretary"},
		},
	}
	cur, err := m.collection.Find(ctx, filter)
	if err != nil {
		return nil, errorx.NewDefaultError("数据库错误: " + err.Error())
	}
	var result []User
	err = cur.All(ctx, &result)
	if err != nil {
		return nil, errorx.NewDefaultError("数据库错误: " + err.Error())
	}
	return result, nil
}

func (m *UserModel) AddUser(id, name, role, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var result User
	find_err := m.collection.FindOne(ctx, bson.D{
		{"id", id},
	}).Decode(&result)
	if find_err == mongo.ErrNoDocuments {
		record := bson.D{
			{"id", id},
			{"password", password},
			{"role", role},
			{"name", name},
		}
		_, err := m.collection.InsertOne(ctx, &record)
		if err != nil {
			return errorx.NewDefaultError("创建失败")
		}
		return nil
	} else if find_err != nil {
		return errorx.NewDefaultError("数据库错误")
	} else {
		return errorx.NewDefaultError("用户ID重复")
	}
}

func (m *UserModel) GetUserBatch(list []User) ([]User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.M{}
	orQuery := bson.D{}
	for _, user := range list {
		orQuery = append(orQuery, bson.E{"id", user.ID})
	}
	filter["$or"] = orQuery
	cur, find_err := m.collection.Find(ctx, filter)
	if find_err == mongo.ErrNoDocuments {
		return []User{}, nil
	} else if find_err != nil {
		return []User{}, errorx.NewDefaultError("数据库错误")
	} else {
		var results []User
		cur.All(ctx, &results)
		return results, nil
	}
}
func (m *UserModel) AddUserBatch(list []User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	users, err := m.GetUserBatch(list)
	if err != nil {
		return err
	}
	if len(users) != 0 {
		return errorx.NewDefaultError("ID重复")
	} else {
		records := []interface{}{}
		for _, user := range list {
			records = append(records, bson.D{
				{"id", user.ID},
				{"password", user.Password},
				{"name", user.Name},
				{"role", user.Role},
			})
		}
		_, err := m.collection.InsertMany(ctx, records)
		if err != nil {
			return errorx.NewDefaultError("数据库错误")
		}
	}
	return nil
}
