package session

import (
	"time"

	"github.com/google/uuid"
)

type Service interface {
	CreateSession() (resp SessionResponse, err error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) CreateSession() (resp SessionResponse, err error) {
	id := uuid.NewString()
	req := Session{
		ID:        id,
		StartedAt: time.Now(),
	}
	if err = s.repo.Create(req); err != nil {
		return resp, err
	}

	resp = SessionResponse{
		Session: id,
	}

	return resp, nil
}
