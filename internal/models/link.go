package model

import (
	"time"

	"github.com/google/uuid"
)


type Link struct {
    ID        uuid.UUID `gorm:"column:link_id;type:char(36);primaryKey;default:(UUID())"` 

    Url       string    `json:"url" gorm:"not null"`
    Code      string    `json:"code" gorm:"unique;not null"`
    Link      string    `json:"link" gorm:"not null"`

    Clicks    uint64     `json:"clicks" gorm:"default:0"`
    LastClick *time.Time `json:"lastClick"`

    CreatedAt time.Time  `gorm:"autoCreateTime"`

    LinkClicks []LinkClick `gorm:"foreignKey:LinkID;constraint:OnDelete:CASCADE" json:"-"`
}


