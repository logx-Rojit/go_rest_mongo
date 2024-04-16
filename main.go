package main

import (
	"fmt"
	"go_rest_mongo/database"
	"go_rest_mongo/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

func StartServer(address string, client *mongo.Client) *ApiStarter {
	return &ApiStarter{
		Addr:   address,
		Client: client,
	}
}

func (s *ApiStarter) Run() {

	r := mux.NewRouter()
	router := r.PathPrefix("/api/v1").Subrouter()

	router.HandleFunc("/users", utils.MakeHTTPHandleFunc(s.GetUsersHandler)).Methods("GET")
	router.HandleFunc("/user/{id}", utils.MakeHTTPHandleFunc(s.GetUserHandler)).Methods("GET")
	router.HandleFunc("/user", utils.MakeHTTPHandleFunc(s.CreateUserHandler)).Methods("POST")
	router.HandleFunc("/user/{id}", utils.MakeHTTPHandleFunc(s.deleteUserById)).Methods("DELETE")

	fmt.Println("Server started at port :8000")
	http.ListenAndServe(s.Addr, r)

}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Panic("Failed to Load env file", err)
	}
	client, err := database.ConnectDb()

	if err != nil {
		log.Panic("Failed to connect to database")
	}
	defer database.Disconnect(client)

	s := StartServer(":8000", client)
	s.Run()
}
