package main

import (
	"context"
	"encoding/json"
	"errors"
	"go_rest_mongo/utils"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type loginType struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	Data         User   `json:"data"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (s *ApiStarter) login(w http.ResponseWriter, r *http.Request) error {

	coll := s.Client.Database("HRM").Collection("user")
	var lg loginType
	var user *User
	var lr loginResponse
	var err error
	var accessToken, refreshToken string
	if err := json.NewDecoder(r.Body).Decode(&lg); err != nil {
		return err
	}

	filter := bson.M{"email": lg.Email}
	err = coll.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return err
	}
	// check if password is valid or not
	err = utils.ComparePassword(lg.Password, user.Password)
	if err != nil {
		return errors.New("please provide valid password cause password did not match")
	}
	jti, _ := utils.GenerateRandomPassword(15, false, true, false)

	// get two token refresh and access token
	accessToken, _ = utils.GenerateToken("access", jti, user.Id.Hex(), time.Now().Add(time.Hour*24).Unix())
	refreshToken, err = utils.GenerateToken("refresh", jti, user.Id.Hex(), time.Now().Add(time.Hour*24).Unix())
	if err != nil {
		return err
	}
	lr.AccessToken = accessToken
	lr.RefreshToken = refreshToken
	lr.Data = *user
	return utils.WriteJSON(w, http.StatusOK, lr)
}

// func (s *ApiStarter) register(w http.ResponseWriter, r *http.Request) error {

// }
