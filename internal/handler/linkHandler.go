package handler

import (
	"net/http"
	"short-url-api/internal/payloads/request"
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
/// @Summary Create a short link
// @Description Generate a short URL from a long URL
// @Tags links
// @Accept json
// @Produce json
// @Param link body request.RequestLink true "URL to shorten"
// @Success 201 {object} dto.LinkDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/links [post]
func (h *LinkHandler) CreateLink(c *gin.Context) {
    var rq request.RequestLink
    if err := c.ShouldBindJSON(&rq); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    linkDTO, err := h.service.CreateLink(c, rq.Url)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, linkDTO)
}