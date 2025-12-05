package service

import (
	"context"
	"short-url-api/internal/payloads/dto"
)

type LinkService interface {
	CreateLink(ctx context.Context,url string) (dto.LinkDTO, error)

	GetByLink(ctx context.Context,link string) (dto.LinkDTO, error)

	GetUrlByCode(ctx context.Context,code string) (string, error)

	GetAllLinks(ctx context.Context, page, limit int) ([]dto.LinkDTO, int64, bool, error)
}