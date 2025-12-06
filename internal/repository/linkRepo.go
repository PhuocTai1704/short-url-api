package repository

import (
	"context"
	"errors"
	model "short-url-api/internal/models"
	"short-url-api/internal/payloads/dto"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SQLLinkRepo struct {
	db *gorm.DB
}

func NewSQLLinkRepo(db *gorm.DB) LinkRepo {
	return &SQLLinkRepo{
		db: db,
	}
}

func (r *SQLLinkRepo) IsCodeExist(ctx context.Context, code string) (bool, error) {
	var link model.Link
	err := r.db.Where("code = ?", code).Take(&link).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *SQLLinkRepo) Create(ctx context.Context, link *model.Link) error {
	if err := r.db.WithContext(ctx).Create(link).Error; err != nil {
		return err
	}
	return nil
}

func (r *SQLLinkRepo) FirstOrCreate(ctx context.Context, link *model.Link) error {
	cond := model.Link{Url: link.Url} // điều kiện tìm

	if err := r.db.WithContext(ctx).Where(&cond).FirstOrCreate(link).Error; err != nil {
		return err
	}
	return nil
}

func (r *SQLLinkRepo) GetByLink(ctx context.Context, urlLink string) (*model.Link, error) {
	var link model.Link

	err := r.db.WithContext(ctx).
		Where("link = ?", urlLink).
		First(&link).Error

	if err != nil {
		return nil, err
	}

	return &link, nil
}

func (r *SQLLinkRepo) GetByCode(ctx context.Context, code string) (*model.Link, error) {
	var link model.Link

	err := r.db.WithContext(ctx).
		Where("code = ?", code).
		First(&link).Error

	if err != nil {
		return nil, err
	}

	return &link, nil
}
func (r *SQLLinkRepo) GetAllLinks(ctx context.Context, page, limit int) ([]model.Link, int64, error) {
	var links []model.Link
	var total int64

	if err := r.db.WithContext(ctx).
		Model(&model.Link{}).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit

	if err := r.db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Find(&links).Error; err != nil {
		return nil, 0, err
	}

	return links, total, nil
}

func (r *SQLLinkRepo) IncreaseClick(ctx context.Context, id uuid.UUID) error {
	r.db.WithContext(ctx)
	return r.db.WithContext(ctx).
		Model(&model.Link{}).
		Where("link_id = ?", id).
		Updates(map[string]any{
			"clicks":     gorm.Expr("clicks + 1"),
			"last_click": time.Now(),
		}).Error
}

func (r *SQLLinkRepo) GetLinkWithStatsByLink(ctx context.Context, urlLink string) (*dto.LinkDTO, error) {
	var dto dto.LinkDTO

	err := r.db.
		WithContext(ctx).
		Table("links AS l").
		Select(`
            l.link_id AS id,
            l.url,
            l.code,
            l.link,
            l.created_at,
            COUNT(c.id) AS clicks,
            MAX(c.clicked_at) AS last_click
        `).
		Joins("LEFT JOIN link_clicks AS c ON c.link_id = l.link_id").
		Where("l.link = ?", urlLink).
		Group("l.link_id").
		Scan(&dto).Error

	if err != nil {
		return nil, err
	}

	return &dto, nil
}
