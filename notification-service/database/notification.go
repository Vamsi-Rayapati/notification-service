package database

import (
	"time"

	"github.com/google/uuid"
)

type NotificationStatus string

const (
	NotificationSending NotificationStatus = "sending"
	NotificationSent    NotificationStatus = "sent"
	NotificationFailed  NotificationStatus = "failed"
)

type NotificationChannel string

type Notification struct {
	ID        uuid.UUID           `gorm:"type:char(36);primaryKey;default:(UUID())"`
	UserID    uuid.UUID           `gorm:"type:char(36);notNull"`
	Channel   NotificationChannel `gorm:"type:enum('email','push');"`
	Status    NotificationStatus  `gorm:"type:enum('sending','sent','failed');"`
	CreatedAt time.Time           `gorm:"autoCreateTime"`
	UpdatedAt time.Time           `gorm:"autoUpdateTime"`
}
