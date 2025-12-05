package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type Link struct {
    ID        uuid.UUID `json:"linkId" gorm:"column:link_id;type:char(36);primaryKey"`

    Url       string    `json:"url" gorm:"unique;not null"`
    Code      string    `json:"code" gorm:"unique;not null"`

    Clicks    uint64     `json:"clicks" gorm:"default:0"`
    LastClick *time.Time `json:"lastClick"`

    CreatedAt time.Time
}

// Auto-generate UUID on create
func (l *Link) BeforeCreate(tx *gorm.DB) (err error) {
    l.ID = uuid.New()
    return
}
