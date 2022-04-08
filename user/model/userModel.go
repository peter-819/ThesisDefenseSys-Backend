package model 

import (
	"context"
	"time"
	
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"TDS-backend/common/mongox"
	"TDS-backend/common/errorx"
)

type User struct {
	ID string `bson:"id"`
	Name string `bson:"name"`
	Password string `bson:"password"`
	Role string `bson:"role"`
}

type IUserModel interface {
	GetUser(id string) (*User, error)
	AddUser(id, name, role, password string) error
}

type UserModel struct {
	collection *mongo.Collection
}

func NewUserModel(d *mongox.Database) IUserModel {
	return &UserModel{
		collection: d.Conn.Collection("UserInfoCollection"),
	}
}

func (m *UserModel) GetUser(id string) (*User, error){
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var result User
	find_err := m.collection.FindOne(ctx,bson.D{
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

func (m *UserModel) AddUser(id, name, role, password string) error{	
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var result User
	find_err := m.collection.FindOne(ctx,bson.D{
		{"id", id},
	}).Decode(&result)
	if find_err == mongo.ErrNoDocuments{
		record := bson.D {
			{"id",id },
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
	} else{
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
func (m *UserModel) AddUserBatch(list []User) error{
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	users, err := m.GetUserBatch(list)
	if err != nil{
		return err
	}
	if len(users) != 0 {
		return errorx.NewDefaultError("ID重复")
	} else{
		records := []interface {}{}
		for _, user := range list{
			records = append(records, bson.D{
				{"id", user.ID},
				{"password", user.Password},
				{"name",user.Name},
				{"role",user.Role},
			})
		}
		_, err := m.collection.InsertMany(ctx, records)
		if err != nil {
			return errorx.NewDefaultError("数据库错误")
		}
	}
	return nil
}