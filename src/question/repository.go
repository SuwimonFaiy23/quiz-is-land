package question

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

type Repository interface {
	GetBySessionID(sessionID string) (resp Question, err error)
	SaveAnswer(req Answer) error
	GetSummary(sessionID string) (int, int, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}



func (r *repository) GetBySessionID(sessionID string) (resp Question, err error) {
	var answered []string
	err = r.db.Model(&Answer{}).
		Where("session_id = ?", sessionID).
		Pluck("question_id", &answered).Error
	if err != nil {
		return resp, err
	}

	tx := r.db.
		Preload("choices").
		Where("id NOT IN ?", answered).
		Order("id ASC").
		First(&resp)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return resp, nil
	}
	if tx.Error != nil {
		return resp, tx.Error
	}

	return resp, nil
}


func (r *repository) SaveAnswer(req Answer) error {
	if err := r.db.Create(&req).Error; err != nil {
		log.Printf("create answer error : %s", err)
		return err
	}
	return nil
}

func (r *repository) GetSummary(sessionID string) (int, int, error) {
	var total int64
	var correct int64

	err := r.db.Model(&Answer{}).
		Where("session_id = ?", sessionID).
		Count(&total).Error
	if err != nil {
		return 0, 0, err
	}

	sub := r.db.Table("answers").
		Select("choice_id").
		Where("session_id = ?", sessionID)

	err = r.db.Model(&Choice{}).
		Where("id IN (?) AND is_correct = true", sub).
		Count(&correct).Error
	if err != nil {
		return 0, 0, err
	}

	return int(total), int(correct), nil
}