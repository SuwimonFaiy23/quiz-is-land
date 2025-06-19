package question

import (
	"time"

	"github.com/SuwimonFaiy23/quiz-is-land/src/session"
)

type Service interface {
	GetQuestion(sessionID string) (resp Question, err error)
	EndSession(sessionID string) error
	SubmitAnswer(req Answer) (err error)
	GetSummary(sessionID string) (resp SummaryResponse, err error)
}

type service struct {
	repo        Repository
	sessionRepo session.Repository
}

func NewService(repo Repository, sessionRepo session.Repository) Service {
	return &service{repo, sessionRepo}
}

func (s *service) GetQuestion(sessionID string) (resp Question, err error) {
	if resp, err = s.repo.GetBySessionID(sessionID); err != nil {
		return resp, err
	}
	return resp, nil
}

func (s *service) EndSession(sessionID string) error {
	return s.sessionRepo.Update(sessionID)
}

func (s *service) SubmitAnswer(req Answer) (err error) {
	reqAns := Answer{
		SessionID:  req.SessionID,
		QuestionID: req.QuestionID,
		ChoiceID:   req.ChoiceID,
		CreatedAt:  time.Now(),
	}
	if err = s.repo.SaveAnswer(reqAns); err != nil {
		return err
	}
	return nil
}

func (s *service) GetSummary(sessionID string) (resp SummaryResponse, err error) {
	total, correct, err := s.repo.GetSummary(sessionID)
	if err != nil {
		return resp, err
	}
	resp = SummaryResponse{
		Total:   total,
		Correct: correct,
	}
	return resp, nil
}
