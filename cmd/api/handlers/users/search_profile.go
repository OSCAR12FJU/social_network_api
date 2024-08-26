package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) SearchProfile(c *gin.Context) {
	paramsId := c.Param("id")

	userProfile, err := h.UserService.Profile(paramsId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "El usuario no existe o hay un error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"user":   userProfile,
	})
}
