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

func (h *LinkHandler) TestHandler(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "OK!",
    })
}
