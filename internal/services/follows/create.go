package follows

import (
	"api_red_social/internal/domain"
	"fmt"
	"log"
	"time"
)

func (s Services) Create(follow domain.Follows) (id interface{}, err error) {

	follow.CreatedAt = time.Now().UTC()

	result, err := s.Repo.Insert(follow)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error creating follower: %w", err)
	}

	return result, nil
}
