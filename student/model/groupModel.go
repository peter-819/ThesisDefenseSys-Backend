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

type Group struct {
	Id        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	DefenseId string             `bson:"defense_id"`
	Members   []string           `bson:"members"`
}

type IGroupModel interface {
	NewGroup(name string) (string, error)
	RemoveGroup(groupId string) error
	// AddMember(groupId string, memberId string) error
	// RemoveMember(groupId string, memberId string) error
	QueryGroup(groupId string) (*Group, error)
	QueryAllGroups() ([]Group, error)
	ModifyGroup(groupId string, newGroup *Group) error
}

type GroupModel struct {
	collection *mongo.Collection
}

func NewGroupModel(d *mongox.Database) IGroupModel {
	return &GroupModel{
		collection: d.Conn.Collection("StudentGroup"),
	}
}

func (m *GroupModel) NewGroup(name string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	group := &Group{}
	group.Id = primitive.NewObjectID()
	group.Name = name
	res, err := m.collection.InsertOne(ctx, group)
	if err != nil {
		return "", errorx.NewDefaultError("新建组失败: " + err.Error())
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (m *GroupModel) RemoveGroup(groupId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	hexId, err := primitive.ObjectIDFromHex(groupId)
	if err != nil {
		return errorx.NewDefaultError("非法组Id: " + err.Error())
	}
	_, err = m.collection.DeleteOne(ctx, bson.M{"_id": hexId})
	if err != nil {
		return errorx.NewDefaultError("删除组失败:" + err.Error())
	}
	return nil
}

// func (m *GroupModel) AddMember(groupId string, memberId string) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
// 	defer cancel()

// 	hexId, err := primitive.ObjectIDFromHex(groupId)
// 	if err != nil {
// 		return errorx.NewDefaultError("非法组Id: " + err.Error())
// 	}
// 	filter := bson.M{"_id": hexId}
// 	update := bson.M{
// 		"$addToSet": bson.M{
// 			"members": memberId,
// 		},
// 	}
// 	res := m.collection.FindOneAndUpdate(ctx, filter, update)
// 	if res.Err() != nil {
// 		return errorx.NewDefaultError("添加成员失败: " + res.Err().Error())
// 	}
// 	return nil
// }
// func (m *GroupModel) RemoveMember(groupId string, memberId string) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
// 	defer cancel()

// 	hexId, err := primitive.ObjectIDFromHex(groupId)
// 	if err != nil {
// 		return errorx.NewDefaultError("非法组Id: " + err.Error())
// 	}
// 	filter := bson.M{"_id": hexId}
// 	update := bson.M{
// 		"$$pull": bson.M{
// 			"members": memberId,
// 		},
// 	}
// 	res := m.collection.FindOneAndUpdate(ctx, filter, update)
// 	if res.Err() != nil {
// 		return errorx.NewDefaultError("删除成员失败: " + res.Err().Error())
// 	}
// 	return nil
// }
func (m *GroupModel) QueryGroup(groupId string) (*Group, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	hexId, err := primitive.ObjectIDFromHex(groupId)
	if err != nil {
		return nil, errorx.NewDefaultError("非法组Id: " + err.Error())
	}

	result := &Group{}
	err = m.collection.FindOne(ctx, bson.M{"_id": hexId}).Decode(result)
	if err != nil {
		return nil, errorx.NewDefaultError("查询组失败: " + err.Error())
	}
	return result, nil
}

func (m *GroupModel) QueryAllGroups() ([]Group, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cur, err := m.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var result []Group
	err = cur.All(ctx, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (m *GroupModel) ModifyGroup(id string, newGroup *Group) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	hexId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errorx.NewDefaultError("非法组Id: " + err.Error())
	}

	filter := bson.M{"_id": hexId}
	update := bson.M{"$set": newGroup}
	res := m.collection.FindOneAndUpdate(ctx, filter, update)
	if res.Err() != nil {
		return errorx.NewDefaultError("修改组失败: " + res.Err().Error())
	}
	return nil
}
