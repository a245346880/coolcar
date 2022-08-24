package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	c := context.Background()
	//构建认证参数
	credential := options.Credential{
		AuthSource: "admin",
		Username:   "root",
		Password:   "123456",
	}
	mc, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://192.168.19.128:27017/coolcar?readPreference=primary&ssl=false").SetAuth(credential))
	if err != nil {
		panic(err)
	}
	defer mc.Disconnect(c)
	col := mc.Database("coolcar").Collection("account")
	findRows(c, col)
}

func findRows(c context.Context, col *mongo.Collection) {
	cur, err := col.Find(c, bson.M{})
	if err != nil {
		panic(err)
	}
	for cur.Next(c) {
		var row struct {
			ID     primitive.ObjectID `bson:"_id"`
			OpenId string             `bson:"open_id"`
		}
		err = cur.Decode(&row)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", row)
	}
}

func insertRows(c context.Context, col *mongo.Collection) {
	res, err := col.InsertMany(c, []interface{}{
		bson.M{
			"open_id": "123",
		},
		bson.M{
			"open_id": "456",
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", res)
}
