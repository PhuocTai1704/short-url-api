package repository

import (
	"context"
	"errors"
	model "short-url-api/internal/models"

	"gorm.io/gorm"
)


type SQLLinkRepo struct {
	db *gorm.DB
}

func NewSQLLinkRepo (db *gorm.DB) LinkRepo {
	return  &SQLLinkRepo{
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
