package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"mongosample.app/client"
)

func main() {
	ctx := context.TODO()
	conn := client.Connect(ctx, "mongodb://localhost:27017")
	client.Save(conn, ctx, "maindb", "movies", bson.M{
		"name":   "Maverick",
		"actor":  "Tom Cruise",
		"budget": "100 million",
	})
	cur, err := client.FindAll(conn, ctx, "maindb", "movies")
	if err != nil {
		panic(err)
		fmt.Println("err", err)
	}
	for cur.Next(ctx) {
		var result bson.D
		cur.Decode(&result)
		for k, v := range result.Map() {
			fmt.Println(k, v)
		}

	}
	cur.Close(ctx)
	conn.Disconnect(ctx)
}
