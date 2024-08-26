package publications

import (
	"api_red_social/internal/domain"
	"fmt"
	"log"
	"time"
)

func (s Services) Create(publi domain.Publication) (id interface{}, err error) {

	publi.CreatedAt = time.Now().UTC()

	result, err := s.Repo.Insert(publi)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error creating publication: %w", err)
	}

	return result, nil
}

// var followCreateParams domain.Follows
// if err := c.BindJSON(&followCreateParams); err != nil {
// 	c.JSON(400, gin.H{"error": err.Error()})
// 	return
// }
// followId, err := h.followService.Create(followCreateParams)
// if err != nil {
// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "oops!"})
// }
// c.JSON(200, gin.H{"follow_id": followId})

// func (service *PublicationService) SavePublication(content, userID string) (*model.Publication, error) {
// 	publication := &model.Publication{
// 		Content:  content,
// 		Follower: userID,
// 	}

// 	return service.Repo.SavePublication(publication)
// }
