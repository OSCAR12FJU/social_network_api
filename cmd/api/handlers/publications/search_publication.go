package publications

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) SearchPublication(c *gin.Context) {
	paramsId := c.Param("id")

	publication, err := h.PublicationService.GetPublicationByID(paramsId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "El usuario no existe o hay un error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":      "success",
		"publication": publication,
	})
}
