package controller

import (
	"api_red_social/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type FollowController struct {
	Service *service.FollowService
}

func NewFollowController(followService *service.FollowService) *FollowController {
	return &FollowController{
		Service: followService,
	}
}

func (c *FollowController) SaveFollow(w http.ResponseWriter, r *http.Request) {
	var params struct {
		Followed string `json:"followed"`
	}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	identity := r.Context().Value("user").(string)

	follow, err := c.Service.SaveFollow(identity, params.Followed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(follow)

}

func (controller *FollowController) DeleteFollow(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(string)
	vars := mux.Vars(r)
	followedID := vars["id"]

	err := controller.Service.DeleteFollow(userID, followedID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Follow eliminado correctamente",
	})

}

func (controller *FollowController) GetFollowing(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(string)
	params := mux.Vars(r)
	page, err := strconv.Atoi(params["page"])
	if err != nil || page < 1 {
		page = 1
	}
	itemsPerPage := 5

	follows, total, err := controller.Service.GetFollowing(userID, page, itemsPerPage)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	following, followers, err := controller.Service.FollowUserIds(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":         "success",
		"message":        "Listado de usuarios que estoy siguiendo",
		"follows":        follows,
		"total":          total,
		"pages":          (total + itemsPerPage - 1) / itemsPerPage,
		"user_following": following,
		"user_follow_me": followers,
	})

}
