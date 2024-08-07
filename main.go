package main

import (
	"api_red_social/controller"
	"api_red_social/database"
	"api_red_social/repository"
	"api_red_social/service"
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	client, err := database.CreateConnection()
	if err != nil {
		log.Fatal("Error al conectar a MongoDB: ", err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Printf("error al hacer ping a MongoDB: %v", err)
		}
	}()

	//Follow
	followRepo := repository.NewFollowRepository(client)
	followService := service.NewFollowService(followRepo)
	followController := controller.NewFollowController(followService)

	//Users
	userRepo := repository.NewUserRepository(client)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	r := mux.NewRouter()

	//Users
	r.HandleFunc("/create-users", userController.CreateUserHandler).Methods("POST")
	r.HandleFunc("/login", userController.LoginHandler).Methods("POST")
	//Follow
	r.HandleFunc("/follow", followController.SaveFollow).Methods("POST")
	r.HandleFunc("/unfollow/{id}", followController.DeleteFollow).Methods("DELETE")
	r.HandleFunc("/following/{page}", followController.GetFollowing).Methods("GET")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8082", nil))

}
