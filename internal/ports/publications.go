package ports

import "api_red_social/internal/domain"

type PublicationService interface {
	Create(publi domain.Publication) (id interface{}, err error)
	GetPublicationByID(publiId string) (id interface{}, err error)
}

type PublicationRepository interface {
	Insert(publi domain.Publication) (id interface{}, err error)
	FindByID(publiId string) (id interface{}, err error)
}
