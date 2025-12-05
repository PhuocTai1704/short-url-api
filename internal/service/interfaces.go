package service

import (
	"short-url-api/internal/payloads/dto"

	"github.com/gin-gonic/gin"
)

type LinkService interface {
	CreateLink(ctx *gin.Context ,url string) (dto.LinkDTO, error)
}