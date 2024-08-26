package users

import (
	"api_red_social/cmd/utils"
	"api_red_social/internal/domain"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s Services) Create(users domain.Users) (id interface{}, err error) {

	users.CreatedAt = time.Now().UTC()
	hashedPassword, err := utils.HashPassword(users.Password)
	if err != nil {
		return primitive.NilObjectID, err
	}
	users.Password = string(hashedPassword)
	users.User_id = primitive.NewObjectID()

	result, err := s.Repo.Insert(users)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error creating follower: %w", err)
	}
	// follow.ID = result.InsertedID.(primitive.ObjectID)
	// return follow, nil
	return result, nil
}

// func (service *UserService) CreateNewUser(users *model.Users) (primitive.ObjectID, error) {
// 	hashedPassword, err := password.Hash(users.Password)
// 	if err != nil {
// 		return primitive.NilObjectID, err
// 	}
// 	users.Password = string(hashedPassword)
// 	users.User_id = primitive.NewObjectID()
// 	users.CreatedAt = time.Now()

// 	return service.Repo.CreateNewUser(users)
// }

// func (c *UserController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
// 	var user model.Users
// 	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	userID, err := c.Service.CreateNewUser(&user)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(map[string]interface{}{
// 		"user_id": userID,
// 	})
// }
