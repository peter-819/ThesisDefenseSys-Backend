package model

import (
	"TDS-backend/common/errorx"
	"TDS-backend/common/mongox"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IRouteModel interface {
	GetMenu(menu interface{}) error
	GetHeader(role string, header interface{}) error
}

type RouteModel struct {
	MenuCollection   *mongo.Collection
	HeaderCollection *mongo.Collection
}

func NewRouteModel(d *mongox.Database) IRouteModel {
	return &RouteModel{
		MenuCollection:   d.Conn.Collection("RoutesMenu"),
		HeaderCollection: d.Conn.Collection("RoutesHeader"),
	}
}

func (m *RouteModel) GetMenu(menu interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	find_err := m.MenuCollection.FindOne(ctx, bson.D{
		{"name", "RouteMenu"},
	}).Decode(menu)
	if find_err != nil {
		return errorx.NewDefaultError("数据库错误: " + find_err.Error())
	}
	return nil
}
func (m *RouteModel) GetHeader(role string, header interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	find_err := m.HeaderCollection.FindOne(ctx, bson.D{
		{"role", role},
	}).Decode(header)
	if find_err != nil {
		return errorx.NewDefaultError("数据库错误: " + find_err.Error())
	}
	return nil
}
