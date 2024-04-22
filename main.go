package main

import (
	"fmt"
	"go_rest_mongo/database"
	"go_rest_mongo/middleware"
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
	newRouter := r.PathPrefix("/api/v1").Subrouter()

	// Login Route
	newRouter.HandleFunc("/login", utils.MakeHTTPHandleFunc(s.login)).Methods("POST")
	newRouter.HandleFunc("/users", middleware.RoleValidatorMiddleware(utils.MakeHTTPHandleFunc(s.GetUsersHandler))).Methods("GET")
	newRouter.HandleFunc("/user/{id}", middleware.RoleValidatorMiddleware(utils.MakeHTTPHandleFunc(s.GetUserHandler))).Methods("GET")
	newRouter.HandleFunc("/user", middleware.RoleValidatorMiddleware(utils.MakeHTTPHandleFunc(s.CreateUserHandler))).Methods("POST")
	newRouter.HandleFunc("/user/{id}", middleware.RoleValidatorMiddleware(utils.MakeHTTPHandleFunc(s.deleteUserById))).Methods("DELETE")

	//Create user login and checkout time
	newRouter.HandleFunc("/checkIn", middleware.RoleValidatorMiddleware(utils.MakeHTTPHandleFunc(s.CheckIn))).Methods("POST")
	newRouter.HandleFunc("/get-list", middleware.RoleValidatorMiddleware(utils.MakeHTTPHandleFunc(s.GetUserCheckIn))).Methods("GET")

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
