package service

import (
	"context"
	"net/http"
	"short-url-api/internal/payloads/dto"

	"github.com/google/uuid"
)

type LinkService interface {
	CreateLink(ctx context.Context, url string, alias string) (dto.LinkDTO, error)

	UpdateLink(ctx context.Context, linkId uuid.UUID, url string, alias string) (dto.LinkDTO, error)

	GetById(ctx context.Context, id uuid.UUID) (dto.LinkDTO, error)

	GetByLink(ctx context.Context, link string) (dto.LinkDTO, error)

	GetUrlByCode(ctx context.Context, code string, req *http.Request) (string, error)

	GetAllLinks(ctx context.Context, page, limit int) ([]dto.LinkDTO, int64, bool, error)

	DeleteById(ctx context.Context, id uuid.UUID) error
}

type LinkClickService interface {
	IncreaseClick(ctx context.Context, linkId uuid.UUID, req *http.Request) error
}
