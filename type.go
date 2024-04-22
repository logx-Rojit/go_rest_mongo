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

type Token struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserId    primitive.ObjectID `json:"userId" bson:"userId"`
	Jti       string             `json:"jti" bson:"jti"`
	ExpireAt  time.Time          `json:"expireAt" bson:"expireAt"`
	TokenType string             `json:"tokenType" bson:"tokenType"`
}

type loggedIN struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserId    primitive.ObjectID `json:"userId" bson:"userId"`
	Date      time.Time          `json:"date" bson:"date"`
	InTime    time.Time          `json:"inTime" bson:"inTime"`
	OutTime   time.Time          `json:"outTime" bson:"outTime"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdateAt  time.Time          `json:"updatedAt" bson:"updatedAt"`
}
