package db

import(
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"context"
)

type Connection struct {
	client *mongo.Client
	ctx *context.Context
	collection *mongo.Collection
}

func Connect(constr string) (*Connection, error) {
	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI(constr)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &Connection { client, &ctx, client.Database("MinecraftServer").Collection("Scan-" + time.Now().Format("01-02-2006 15:04:05")) }, nil
}

func Add[T any](con *Connection, entry *T) error {
	_, err := con.collection.InsertOne(*con.ctx, entry)
	return err
} 

func Close(con *Connection) {
	con.client.Disconnect(*con.ctx);
}