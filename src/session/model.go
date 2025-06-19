package session

import "time"

type Session struct {
	ID        string     `gorm:"primaryKey;column:id"`
	StartedAt time.Time  `gorm:"column:started_at"`
	EndedAt   *time.Time `gorm:"column:ended_at"`
}

type SessionResponse struct {
	Session string `json:"session"`
}
