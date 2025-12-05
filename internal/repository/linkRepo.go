package repository

import (
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
func (r *SQLLinkRepo) Create(link *model.Link) error {
	if err := r.db.Create(link).Error; err != nil {
		return err
	}
	return nil
}

func (r *SQLLinkRepo) FirstOrCreate(link *model.Link) error {
    return r.db.Where("url = ?", link.Url).FirstOrCreate(link).Error
}