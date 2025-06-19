package session

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	Create(req Session) error
	Update(sessionID string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(req Session) error {
	if err := r.db.Create(&req).Error; err != nil {
		log.Printf("create session error : %s", err)
		return err
	}
	return nil
}

func (r *repository) Update(sessionID string) error {
	if err := r.db.Debug().Model(&Session{}).Where("id = ?", sessionID).Update("ended_at", time.Now()).Error; err != nil {
		log.Printf("create session error : %s", err)
		return err
	}
	return nil
}
