package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go_rest_mongo/utils"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *ApiStarter) GetUsersHandler(w http.ResponseWriter, r *http.Request) error {
	// query := r.URL.Query()
	// fmt.Println(query.Get("page"), query.Get("pageSize"))
	var user []*User
	coll := s.client.Database("HRM").Collection("user")
	data, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		return err
	}
	if err := data.All(context.Background(), &user); err != nil {
		return err
	}
	return utils.WriteJSON(w, http.StatusOK, user)
}

func (s *ApiStarter) GetUserHandler(w http.ResponseWriter, r *http.Request) error {
	var user *User
	coll := s.client.Database("HRM").Collection("user")
	id := getId(r)
	fmt.Println(id)
	err := coll.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return err
	}
	return utils.WriteJSON(w, http.StatusOK, user)
}

func (s *ApiStarter) CreateUserHandler(w http.ResponseWriter, r *http.Request) error {
	coll := s.client.Database("HRM").Collection("user")
	var user *User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return err
	}
	_, err := coll.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return utils.WriteJSON(w, http.StatusOK, "User created successfully!!!")
}

func getId(r *http.Request) string {
	id := mux.Vars(r)["id"]
	return id
}
