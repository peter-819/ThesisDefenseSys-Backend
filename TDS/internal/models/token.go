package models
import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"TDS-backend/TDS/internal/types"
	"fmt"
	"context"
    "time"
)
func GenerateToken(username string) string{
	return username + "_token"
}

func QueryUserInfo(Database *mongo.Database, token string) (*types.DbUserInfo, error) {
	fmt.Println("querying token: ", token)
	collection := Database.Collection("UserInfoCollection")
	
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	var result types.DbUserInfo

	find_err := collection.FindOne(ctx,bson.D{
		{"token", token},
	}).Decode(&result)
	return &result, find_err
}