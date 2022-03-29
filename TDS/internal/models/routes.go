package models
import (
	"fmt"
	"context"
    "time"

	"go.mongodb.org/mongo-driver/mongo"
	"TDS-backend/TDS/internal/types"
	"go.mongodb.org/mongo-driver/bson"
)

func QueryRoutesMenu(Database *mongo.Database) (*types.RouteMenuReply, error) {
	fmt.Println("querying Route Menu: ")
	collection := Database.Collection("RoutesMenuCollection")
	
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	var result types.RouteMenuReply

	find_err := collection.FindOne(ctx,bson.D{
		{"name", "RouteMenu"},
	}).Decode(&result)
	fmt.Println("router menu result: ", result)
	return &result, find_err
}

func QueryRouteHeader(Database *mongo.Database, role string) (*types.HeaderReply, error) {
	fmt.Println("querying Route Header: ")
	collection := Database.Collection("RoutesHeaderCollection")
	
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	var result types.HeaderReply

	find_err := collection.FindOne(ctx,bson.D{
		{"role", role},
	}).Decode(&result)
	fmt.Println("header result: ", result)
	return &result, find_err
}