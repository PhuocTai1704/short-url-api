package dto

import (
	"time"

	"github.com/google/uuid"
)

type LinkDTO struct {
	ID uuid.UUID `json:"linkId"`

	Name   string `json:"name"`
	Url    string `json:"url"`
	Code string `json:"code"`

	Clicks    uint64     `json:"clicks"`
	LastClick *time.Time `json:"lastClick"`

	CreatedAt time.Time `json:"createdAt"`
}