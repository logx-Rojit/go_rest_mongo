package main

import (
	"context"
	"encoding/json"
	"go_rest_mongo/utils"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *ApiStarter) CheckIn(w http.ResponseWriter, r *http.Request) error {
	coll := s.Client.Database("HRM").Collection("loggedIns")
	var checkInTime *loggedIN
	if err := json.NewDecoder(r.Body).Decode(&checkInTime); err != nil {
		return err
	}

	data, err := coll.InsertOne(context.Background(), checkInTime)

	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, data)

}

func (s *ApiStarter) GetUserCheckIn(w http.ResponseWriter, r *http.Request) error {
	coll := s.Client.Database("HRM").Collection("loggedIns")
	var allData []*loggedIN

	data, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		return err
	}
	if err := data.All(context.Background(), &allData); err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, allData)
}

func (s *ApiStarter) UpdateCheckOutTime(w http.ResponseWriter, r *http.Request) error {
	var data *loggedIN
	coll := s.Client.Database("HRM").Collection("loggedIns")
	id := ""
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil
	}
	sc := bson.D{{"_id", userId}}
	uc := bson.D{{"updatedAt", time.Now()}}
	opts := options.FindOneAndUpdate().
		SetReturnDocument(options.After)

	if err := coll.FindOneAndUpdate(context.Background(), sc, uc, opts).Decode(&data); err != nil {
		return err
	}
	res, _ := bson.MarshalExtJSON(data, false, false)
	return utils.WriteJSON(w, http.StatusOK, res)

}
