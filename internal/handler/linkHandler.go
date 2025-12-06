package handler

import (
	"net/http"
	"short-url-api/internal/payloads/request"
	"short-url-api/internal/payloads/response"
	"short-url-api/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LinkHandler struct {
	service service.LinkService
}

func NewLinkHandler(service service.LinkService) *LinkHandler {
	return &LinkHandler{
		service: service,
	}
}

// / @Summary Create a short link
// @Description Generate a short URL from a long URL
// @Tags links
// @Accept json
// @Produce json
// @Param link body request.RequestLink true "URL to shorten"
// @Success 201 {object} dto.LinkDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /links [post]
func (h *LinkHandler) CreateLink(c *gin.Context) {
	var rq request.RequestLink
	if err := c.ShouldBindJSON(&rq); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	linkDTO, err := h.service.CreateLink(c.Request.Context(), rq.Url, rq.Alias)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, linkDTO)
}

// GetByLink godoc
// @Summary      Lấy thông tin link theo linkID
// @Description  Trả về thông tin chi tiết của link dựa vào linkID
// @Tags         links
// @Accept       json
// @Produce      json
// @Param        id   path     string  true   "Link ID"
// @Success      200   {object}  dto.LinkDTO
// @Failure      400   {object}  response.ErrorResponse
// @Failure      404   {object}  response.ErrorResponse
// @Router       /links/{id} [get]
func (h *LinkHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	linkId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "ID không hợp lệ",
		})
		return
	}
	linkDTO, err := h.service.GetById(c.Request.Context(), linkId)
	if err != nil {
		c.JSON(http.StatusNotFound, response.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, linkDTO)
}

// GetByLink godoc
// @Summary      Lấy thông tin link theo short Link
// @Description  Trả về thông tin chi tiết của link dựa vào short Link
// @Tags         links
// @Accept       json
// @Produce      json
// @Param        url   query     string  true   "Short Link"
// @Success      200   {object}  dto.LinkDTO
// @Failure      400   {object}  response.ErrorResponse
// @Failure      404   {object}  response.ErrorResponse
// @Router       /links/link [get]
func (h *LinkHandler) GetByLink(c *gin.Context) {
	urlLink := c.Query("url") // lấy ?url=

	if urlLink == "" {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Message: "url query param is required",
		})
		return
	}

	linkDTO, err := h.service.GetByLink(c.Request.Context(), urlLink)
	if err != nil {
		c.JSON(http.StatusNotFound, response.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, linkDTO)
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
// @Router       /links [get]
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

// Redirect godoc
// @Summary      Chuyển hướng đến URL gốc bằng short link
// @Description  Nhận mã short link và redirect người dùng đến URL gốc
// @Tags         redirect
// @Param        code   path      string  true   "Short link code"
// @Success      302    "Redirect đến URL gốc"
// @Failure      404    {object}  response.ErrorResponse "Short link không tồn tại"
// @Router       /{code} [get]
func (h *LinkHandler) Redirect(ctx *gin.Context) {
	code := ctx.Param("code")

	url, err := h.service.GetUrlByCode(ctx.Request.Context(), code, ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{
			Message: "Short link not found",
		})
		return
	}

	// Redirect 301 hoặc 302
	ctx.Redirect(http.StatusFound, url)
}
