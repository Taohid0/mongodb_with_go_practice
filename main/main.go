package main

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"log"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {
	client, err := mongo.Connect(context.TODO(), "mongodb://localhost:27017")

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	collection := client.Database("test").Collection("trainers")

	//ash := Trainer{"Ash",10,"Pallet Town"}
	//misty := Trainer{"Misty", 10, "Cerulean City"}
	//brock := Trainer{"Brock", 15, "Power City"}

	//inserResult, err := collection.InsertOne(context.TODO(),ash)
	//
	//if err!=nil{
	//	log.Fatal(err)
	//}
	//fmt.Println(inserResult)

	//trainers := []interface{}{misty, brock}
	//
	//insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(insertManyResult)

	//filter := bson.D{{"name", "Misty"}}

	//update := bson.D{
	//	{"$inc",bson.D{
	//		{"age",1},
	//	}},
	//}
	//
	//updateResult, err := collection.UpdateOne(context.TODO(),filter,update)
	//
	//if err!=nil{
	//	fmt.Println(nil)
	//}
	//fmt.Println(updateResult)

	//var result Trainer
	//err = collection.FindOne(context.TODO(), filter).Decode(&result)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(result)

	findOptions := options.Find()
	findOptions.SetLimit(2)
	var results []*Trainer

	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal("Find error :", err)
	}
	for cur.Next(context.TODO()) {
		var elem Trainer
		err := cur.Decode(&elem)

		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	err = cur.Close(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	for index,value := range results{
		fmt.Println(index,*value)
	}
	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

}
