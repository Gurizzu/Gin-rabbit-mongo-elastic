package database

import (
	"context"
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoUtil struct {
	srv            string
	dbName         string
	collectionName string
	ctx            context.Context
}

func NewMongodbUtil(collectionName string) *MongoUtil {
	return &MongoUtil{
		srv:            "mongodb://localhost:27017",
		dbName:         "go-base",
		collectionName: collectionName,
		ctx:            context.Background(),
	}
}

func (this MongoUtil) Connect() (client *mongo.Client, err error) {
	clientOptions := options.Client()
	clientOptions.ApplyURI(this.srv)

	client, err = mongo.Connect(this.ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		return
	}

	return
}

func (this MongoUtil) Disconnect(client *mongo.Client) {
	if client == nil {
		return
	}
	if err := client.Disconnect(this.ctx); err != nil {
		log.Println(err)
		return
	}

}

func (this MongoUtil) Upsert(data interface{}) (resp interface{}) {

	client, err := this.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer this.Disconnect(client)
	col := client.Database(this.dbName).Collection(this.collectionName)

	_, err = col.InsertOne(this.ctx, data)
	if err != nil {
		log.Fatal(err)
	}

	return data

}

func (this MongoUtil) BaseFindOne(filter bson.M, pointerDecodeTo interface{}) (err error) {

	client, err := this.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer this.Disconnect(client)
	col := client.Database(this.dbName).Collection(this.collectionName)

	res := col.FindOne(this.ctx, filter)

	if res.Err() != nil {
		filterAsJson, _ := json.Marshal(filter)
		log.Println(err, this.collectionName, string(filterAsJson))
		err = errors.New("Data not found")
		return
	}

	if err := res.Decode(pointerDecodeTo); err != nil {
		log.Println(err)
	}
	return
}

func (this MongoUtil) FindOne(key string, value string, pointerDecodeTo interface{}) (err error) {
	objId, err := primitive.ObjectIDFromHex(value)
	if err != nil {
		log.Println(err)
	}
	return this.BaseFindOne(bson.M{key: objId}, pointerDecodeTo)
}
