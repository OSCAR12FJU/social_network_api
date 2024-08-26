package publications

import "api_red_social/internal/ports"

type Handler struct {
	PublicationService ports.PublicationService
}
