package client

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(ctx context.Context, uri string) *mongo.Client {
	conn, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println("error in connecting to mongodb")
		return nil
	}
	return conn
}

func Save(conn *mongo.Client, ctx context.Context, db string, collection string, data primitive.M) {
	conn.Database(db).Collection(collection).InsertOne(ctx, data)
}

func FindAll(conn *mongo.Client, ctx context.Context, db string, collection string) (*mongo.Cursor, error) {
	cursor, err := conn.Database(db).Collection(collection).Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	return cursor, nil
}
