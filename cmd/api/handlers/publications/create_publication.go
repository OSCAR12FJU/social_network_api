package publications

import (
	"api_red_social/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) CreatePublication(c *gin.Context) {

	var params map[string]interface{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	text, ok := params["text"].(string)
	if !ok || text == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Datos Invalidos!"})
		return
	}

	userID := c.GetHeader("X-User-Id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	var publiCreateParams domain.Publication

	if err := c.BindJSON(&publiCreateParams); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	publiId, err := h.PublicationService.Create(publiCreateParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "oops!"})
	}
	c.JSON(200, gin.H{"Publication_id": publiId})

}

// func (controller *PublicationController) SavePublication(w http.ResponseWriter, r *http.Request) {
// 	var params map[string]interface{}
// 	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
// 		http.Error(w, "Datos inválidos", http.StatusBadRequest)
// 		return
// 	}

// 	text, ok := params["text"].(string)
// 	if !ok || text == "" {
// 		http.Error(w, "Datos inválidos", http.StatusBadRequest)
// 		return
// 	}

// 	userID := r.Header.Get("X-User-Id")
// 	if userID == "" {
// 		http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 		return
// 	}

// 	publication, err := controller.Service.SavePublication(text, userID)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	payload := map[string]interface{}{
// 		"status":            "success",
// 		"message":           "Publicación guardada",
// 		"publicationStored": publication,
// 	}

// 	response, _ := json.Marshal(payload)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(response)

// }
