package repository

import (
	"context"
	model "short-url-api/internal/models"
	"short-url-api/internal/payloads/dto"

	"github.com/google/uuid"
)

type LinkRepo interface {
	IsCodeExist(ctx context.Context, code string) (bool, error)

	Create(ctx context.Context, link *model.Link) error

	FirstOrCreate(ctx context.Context, link *model.Link) error

	GetByLink(ctx context.Context, url string) (*model.Link, error)

	GetByCode(ctx context.Context, code string) (*model.Link, error)

	GetAllLinks(ctx context.Context, page, limit int) ([]model.Link, int64, error)

	IncreaseClick(ctx context.Context, id uuid.UUID) error

	GetLinkWithStatsByLink(ctx context.Context,urlLink string) (*dto.LinkDTO, error)
}

type LinkClickRepo interface {
	Create(ctx context.Context, linkClick *model.LinkClick) error
}
