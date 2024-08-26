package follows

import (
	"api_red_social/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) CreateFollow(c *gin.Context) {
	var followCreateParams domain.Follows
	if err := c.BindJSON(&followCreateParams); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	followId, err := h.FollowService.Create(followCreateParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "oops!"})
	}
	c.JSON(200, gin.H{"follow_id": followId})

}

// func (repo *FollowRepository) SaveFollow(follow  *model.Follow) (*model.Follow, error) {

// }
