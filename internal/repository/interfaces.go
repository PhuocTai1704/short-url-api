package repository

import (
	"context"
	model "short-url-api/internal/models"
)

type LinkRepo interface {

	IsCodeExist(ctx context.Context, code string) (bool, error)

	Create(ctx context.Context,link *model.Link) error
	
	FirstOrCreate(ctx context.Context,link *model.Link) error

	GetByLink(ctx context.Context, url string) (*model.Link, error)
	
	GetAllLinks(ctx context.Context, page, limit int) ([]model.Link, int64, error) 
}
