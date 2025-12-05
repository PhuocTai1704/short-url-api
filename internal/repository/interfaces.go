package repository

import (
	"context"
	model "short-url-api/internal/models"
)

type LinkRepo interface {
	Create(ctx context.Context,link *model.Link) error
	
	FirstOrCreate(ctx context.Context,link *model.Link) error

	GetAllLinks(ctx context.Context, page, limit int) ([]model.Link, int64, error) 
}
