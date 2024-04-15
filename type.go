package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ApiStarter struct {
	Addr   string
	client *mongo.Client
}

type User struct {
	Id        primitive.ObjectID  `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string              `json:"name" bson:"name"`
	Email     string              `json:"email" bson:"email,unique"`
	Dob       int                 `json:"dob" bson:"dob"`
	Contact   string              `json:"contact" bson:"contact"`
	CreatedAt primitive.Timestamp `json:"createdAt" bson:"createdAt"`
	UpdatedAt primitive.Timestamp `json:"updatedAt" bson:"updatedAt"`
}
