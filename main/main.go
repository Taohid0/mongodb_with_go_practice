package main

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
)

type Trainer struct {
	Name string
	Age int
	City string
}

func main()  {
	client, err := mongo.Connect(context.TODO(),"mongodb://localhost:27017")

	if err!=nil{
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(),nil)

	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	collection := client.Database("test").Collection("trainers")


	ash := Trainer{"Ash",10,"Pallet Town"}
	misty := Trainer{"Misty",10,"Cerulean City"}
	brock := Trainer{"Brock",15,"Power City"}



	err = client.Disconnect(context.TODO())

	if err!=nil{
		log.Fatal(err)
	}

}