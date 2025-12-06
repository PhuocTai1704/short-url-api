package repository

import (
	"context"
	model "short-url-api/internal/models"
	"time"

	"gorm.io/gorm"
)

type SQLLinkClickRepo struct {
	db *gorm.DB
}

func NewSQLLinkClickRepo(db *gorm.DB) LinkClickRepo {
	return &SQLLinkClickRepo{
		db: db,
	}
}


func (r *SQLLinkClickRepo) Create(ctx context.Context, click *model.LinkClick) error {
    // Tự động set thời gian click nếu chưa có
    if click.ClickedAt.IsZero() {
        click.ClickedAt = time.Now()
    }

    // Lưu vào DB
    if err := r.db.WithContext(ctx).Create(click).Error; err != nil {
        return err
    }

    return nil
}
