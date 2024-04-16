package main

import (
	"context"
	"encoding/json"
	"go_rest_mongo/utils"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *ApiStarter) GetUsersHandler(w http.ResponseWriter, r *http.Request) error {
	// query := r.URL.Query()
	// fmt.Println(query.Get("page"), query.Get("pageSize"))
	var user []*User
	coll := s.Client.Database("HRM").Collection("user")
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
	var user User
	var err error
	coll := s.Client.Database("HRM").Collection("user")
	id := getId(r)
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	err = coll.FindOne(context.Background(), bson.M{"_id": userId}).Decode(&user)
	if err != nil {
		return err
	}
	return utils.WriteJSON(w, http.StatusOK, user)
}

func (s *ApiStarter) CreateUserHandler(w http.ResponseWriter, r *http.Request) error {
	var user *User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return err
	}
	coll := s.Client.Database("HRM").Collection("user")
	HashPassword, password := utils.GenerateRandomPassword(16, true, true, true)

	err := utils.SendEmail(user.Email, "User Created successfully!!!", "This is new Email and thanks for creating user using this application and your password "+password+" thanks")

	user.Password = HashPassword
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if err != nil {
		return utils.WriteJSON(w, http.StatusFailedDependency, err)
	}
	_, err = coll.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return utils.WriteJSON(w, http.StatusOK, user)
}

func (s *ApiStarter) deleteUserById(w http.ResponseWriter, r *http.Request) error {
	id := getId(r)

	coll := s.Client.Database("HRM").Collection("user")
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = coll.DeleteOne(context.Background(), bson.M{"_id": userId})
	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, "User Deleted Successfully!!!")
}

func getId(r *http.Request) string {
	id := mux.Vars(r)["id"]
	return id
}
