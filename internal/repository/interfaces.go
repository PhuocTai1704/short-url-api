package repository

import (
	"context"
	model "short-url-api/internal/models"
	"short-url-api/internal/payloads/dto"
	"time"

	"github.com/google/uuid"
)

type LinkRepo interface {
	IsCodeExist(ctx context.Context, code string) (bool, error)

	IsCodeExistNotId(ctx context.Context, linkId uuid.UUID, code string) (bool, error)

	Create(ctx context.Context, link *model.Link) error

	Update(ctx context.Context, link *model.Link) error

	FirstOrCreate(ctx context.Context, link *model.Link) error

	GetByLink(ctx context.Context, url string) (*model.Link, error)

	GetByCode(ctx context.Context, code string) (*model.Link, error)

	ExistsByCode(ctx context.Context, code string) (bool, error)

	GetAllLinks(ctx context.Context, page, limit int, startDate, endDate *time.Time) ([]model.Link, int64, error)

	GetLinkWithStatsByLink(ctx context.Context, urlLink string) (*dto.LinkDTO, error)

	GetLinkWithStatsById(ctx context.Context, id uuid.UUID) (*dto.LinkDTO, error)

	DeleteById(ctx context.Context, id uuid.UUID) error
}

type LinkClickRepo interface {
	Create(ctx context.Context, linkClick *model.LinkClick) error
}
