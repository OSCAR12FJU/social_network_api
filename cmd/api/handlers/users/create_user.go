package users

import (
	"api_red_social/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) CreateUser(c *gin.Context) {
	var userCreateParams domain.Users

	if err := c.BindJSON(&userCreateParams); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	userID, err := h.UserService.Create(userCreateParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "oops!"})
	}
	c.JSON(200, gin.H{"User_id": userID})

}

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
