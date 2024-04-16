package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ApiStarter struct {
	Addr   string
	Client *mongo.Client
}

type User struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	Email     string             `json:"email" bson:"email,unique"`
	Password  string             `json:"password" bson:"password"`
	Dob       int                `json:"dob" bson:"dob"`
	Contact   string             `json:"contact" bson:"contact"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type EmailTemplate struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Slug      string             `json:"slug" bson:"slug"`
	Body      string             `json:"body" bson:"body"`
	Subject   string             `json:"subject" bson:"subject"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}
