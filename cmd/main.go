package main

import (
	followHandler "api_red_social/cmd/api/handlers/follows"
	publiHandler "api_red_social/cmd/api/handlers/publications"
	userHandler "api_red_social/cmd/api/handlers/users"
	"api_red_social/internal/repositories/mongo"
	followRepo "api_red_social/internal/repositories/mongo/follows"
	publiRepo "api_red_social/internal/repositories/mongo/publications"
	userRepo "api_red_social/internal/repositories/mongo/users"
	followService "api_red_social/internal/services/follows"
	publiService "api_red_social/internal/services/publications"
	userService "api_red_social/internal/services/users"
	"context"
	"fmt"

	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUserName := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	if dbUserName == "" {
		log.Fatalf("DB_USER not set in .env file")
	} else {
		fmt.Println("DB_USER:", dbUserName)
	}

	ginEngine := gin.Default()

	client, err := mongo.ConnectClient(os.Getenv("MONGO_URI"))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Printf("error al hacer ping a MongoDB: %v", err)
		}
	}()
	//Users
	userRepos := userRepo.Repository{}
	userSrv := userService.Services{
		Repo: userRepos}
	userHandle := userHandler.Handler{
		UserService: userSrv}
	// Publication
	publiRepos := publiRepo.Repository{}
	publiSrv := publiService.Services{
		Repo: publiRepos}
	publiHandle := publiHandler.Handler{
		PublicationService: publiSrv}
	// Followers
	followRepos := followRepo.Repository{}
	followSrv := followService.Services{
		Repo: followRepos}
	followHandle := followHandler.Handler{
		FollowService: followSrv}

	ginEngine.POST("/create-user", userHandle.CreateUser)
	ginEngine.GET("/get-user", userHandle.SearchProfile)

	ginEngine.POST("/create-publi", publiHandle.CreatePublication)

	ginEngine.POST("/create-follow", followHandle.CreateFollow)

	log.Fatalln(ginEngine.Run(":8082"))

}
