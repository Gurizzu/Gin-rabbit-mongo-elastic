package services

import (
	"context"
	db "gin_es-rabbit/database"
	"gin_es-rabbit/database/rabbitmq/msg"
	"gin_es-rabbit/models"
	"log"
)

type BookServices struct {
	collectionName string
	ctx            context.Context
	dbUtil         *db.MongoUtil
	rabbitUtil     *msg.RabbitUtil
}

func NewBookServices() *BookServices {
	this := &BookServices{
		collectionName: "books",
		ctx:            context.Background(),
	}
	this.dbUtil = db.NewMongodbUtil(this.collectionName)
	this.rabbitUtil = msg.NewRabbitUtil(this.collectionName)

	return this
}

func (this *BookServices) InsertOne(param models.Book) (resp models.Response) {
	data := this.dbUtil.Upsert(param)
	err := this.rabbitUtil.Send(param)
	if err != nil {
		log.Println(err)
	}
	resp.Data = data

	return
}

func (this *BookServices) GetOne(key string, value string) (res models.Response) {
	var resp models.Book
	err := this.dbUtil.FindOne(key, value, &resp)
	if err != nil {
		log.Println(err)
		return
	}
	res.Data = resp
	return
}
