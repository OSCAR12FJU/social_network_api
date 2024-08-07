package controller

import (
	"api_red_social/model"
	"api_red_social/service"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service *service.UserService
}

func (controller *UserController) Profile(c *gin.Context) {
	id := c.Param("id")

	userProfile, err := controller.Service.GetUserProfile(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "El usuario no existe o hay un error",
		})
		return
	}

	// followInfo, err := controller.Service.GetFollowInfo(c.GetString("userId"), id)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"status":  "error",
	// 		"message": "Error al obtener información de seguimiento",
	// 	})
	// 	return
	// }
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"user":   userProfile,
		// "following":  followInfo.Following,
		// "follower":   followInfo.Follower,
	})
}

func (c *UserController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var params struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, "Faltan datos por enviar", http.StatusBadRequest)
		return
	}

	token, user, err := c.Service.Login(params.Email, params.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response := struct {
		Status  string      `json:"status"`
		Message string      `json:"message"`
		User    interface{} `json:"user"`
		Token   string      `json:"token"`
	}{
		Status:  "success",
		Message: "Te has identificado correctamente",
		User: map[string]interface{}{
			"id":   user.User_id,
			"name": user.Name,
			"nick": user.Nick,
		},
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{Service: userService}
}

func (c *UserController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user model.Users
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, err := c.Service.CreateNewUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id": userID,
	})
}
