package handler

import (
	"net/http"
	"short-url-api/internal/payloads/request"
	"short-url-api/internal/payloads/response"
	"short-url-api/internal/service"
	"strconv"

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

    linkDTO, err := h.service.CreateLink(c.Request.Context(), rq.Url)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, linkDTO)
}
// GetAllLinks godoc
// @Summary      Lấy danh sách link
// @Description  Lấy danh sách link theo phân trang
// @Tags         links
// @Accept       json
// @Produce      json
// @Param        page   query     int     false  "Trang hiện tại"      default(1)
// @Param        limit  query     int     false  "Số item mỗi trang"   default(10)
// @Success      200    {object}  response.LinkResponse
// @Failure      400    {object}  response.ErrorResponse
// @Failure      500    {object}  response.ErrorResponse
// @Router       /api/links [get]
func (h *LinkHandler) GetAllLinks(c *gin.Context) {
    ctx := c.Request.Context()

    page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
    if err != nil || page <= 0 {
        c.JSON(http.StatusBadRequest, response.ErrorResponse{
            Message: "invalid page",
        })
        return
    }

    limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
    if err != nil || limit <= 0 {
        c.JSON(http.StatusBadRequest, response.ErrorResponse{
            Message: "invalid limit",
        })
        return
    }

    links, total, isLast, err := h.service.GetAllLinks(ctx, page, limit)
    if err != nil {
        c.JSON(http.StatusInternalServerError, response.ErrorResponse{
            Message: err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, response.LinkResponse{
        Data:       links,
        Total:      total,
        Page:       page,
        Limit:      limit,
        IsLastPage: isLast,
    })
}
