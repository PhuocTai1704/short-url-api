package model

import (
	"time"

	"github.com/google/uuid"
)

type LinkClick struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	LinkID    uuid.UUID `gorm:"type:char(36);not null"`
	ClickedAt time.Time `gorm:"not null"`  // thời gian click
	IP        string    `gorm:"size:45"`   // IPv4/IPv6
	Referrer  string    `gorm:"size:2083"` // URL giới thiệu
	Country   string    `gorm:"size:255"`
	City      string    `gorm:"size:255"`
	Device    string    `gorm:"size:50"`  // mobile, desktop, tablet
	Language  string    `gorm:"size:255"` // ngôn ngữ trình duyệt
	OS        string    `gorm:"size:100"`
	Browser   string    `gorm:"size:255"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
