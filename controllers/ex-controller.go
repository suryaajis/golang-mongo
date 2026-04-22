package controllers

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOne(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {

	// select database and collection ith Client.Database method
	// and Database.Collection method
	collection := client.Database(dataBase).Collection(col)

	// InsertOne accept two argument of type Context
	// and of empty interface
	result, err := collection.InsertOne(ctx, doc)
	return result, err
}

func InsertMany(client *mongo.Client, ctx context.Context, dataBase, col string, docs []interface{}) (*mongo.InsertManyResult, error) {

	// select database and collection ith Client.Database
	// method and Database.Collection method
	collection := client.Database(dataBase).Collection(col)

	// InsertMany accept two argument of type Context
	// and of empty interface
	result, err := collection.InsertMany(ctx, docs)
	return result, err
}

// var document interface{}

// 	document = bson.D{
// 		{"rollNo", 175},
// 		{"maths", 80},
// 		{"science", 90},
// 		{"computer", 95},
// 	}

// var documents []interface{}
// documents = []interface{}{
// 	bson.D{
// 		{"rollNo", 153},
// 		{"maths", 65},
// 		{"science", 59},
// 		{"computer", 55},
// 	},
// 	bson.D{
// 		{"rollNo", 162},
// 		{"maths", 86},
// 		{"science", 80},
// 		{"computer", 69},
// 	},
// }

// result, err := controllers.InsertOne(client, ctx, "bookshelf", "marks", document)
// if err != nil {
// 	panic(err)
// }

// fmt.Println(result)

// manyResult, err := controllers.InsertMany(client, ctx, "bookshelf", "marks", documents)
// if err != nil {
// 	panic(err)
// }
// fmt.Println(manyResult)
