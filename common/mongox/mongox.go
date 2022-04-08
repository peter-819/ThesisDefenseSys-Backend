package mongox

import (
	"context"
    "time"
    "log"

	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)
type Database struct {
	Conn *mongo.Database
}
type DataInfo struct{
	DataSource string
	DatabaseName string
}
func ConnectMongoDB(Info DataInfo) *Database {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(Info.DataSource).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
        panic(err)
	}
	return &Database{
		Conn: client.Database(Info.DatabaseName),
	}
}
