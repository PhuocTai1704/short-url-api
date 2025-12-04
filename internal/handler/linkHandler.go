package handler

import (
	"short-url-api/internal/service"

	"github.com/gin-gonic/gin"
)

type LinkHandler struct {
	service service.LinkService
}

func NewLinkHandler(service service.LinkService) *LinkHandler {
	return &LinkHandler{
		service: service,
	}
}

// @Summary Test API
// @Description Test handler to check server
// @Tags links
// @Produce json
// @Success 200 {object} map[string]string
// @Router /api/links/test [get]
func (h *LinkHandler) TestHandler(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "OK!",
    })
}
