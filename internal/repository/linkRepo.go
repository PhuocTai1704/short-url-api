package repository

import "gorm.io/gorm"


type SQLLinkRepo struct {
	db *gorm.DB
}

func NewSQLLinkRepo (db *gorm.DB) LinkRepo {
	return  &SQLLinkRepo{
		db: db,
	}
}