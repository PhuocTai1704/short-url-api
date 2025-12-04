package dto

import (
	"time"

	"github.com/google/uuid"
)

type LinkDTO struct {
	ID uuid.UUID `json:"linkId"`

	Name   string `json:"name"`
	Url    string `json:"url"`
	UrlNew string `json:"urlNew"`

	Clicks    uint64     `json:"clicks"`
	LastClick *time.Time `json:"lastClick"`

	IPAddress string `json:"ipAddress"`
	UserAgent string `json:"userAgent"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}