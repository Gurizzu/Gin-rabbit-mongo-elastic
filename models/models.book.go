package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID          primitive.ObjectID `bson:"_id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Created_At  int64              `json:"created_at"`
	Updated_At  int64              `json:"updated_at"`
	Price       float64            `json:"price"`
}

type BookInput struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
