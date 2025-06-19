package question

import "time"

type Question struct {
	ID     int      `gorm:"primaryKey;column:id" json:"session"`
	Text   string   `gorm:"column:text" json:"text"`
	Choice []Choice `gorm:"foreignKey:QuestionID" json:"choice"`
}

type Choice struct {
	ID         int    `gorm:"primaryKey;column:id"`
	QuestionID string `gorm:"column:question_id"`
	Text       string `gorm:"column:text"`
	IsCorrect  string `gorm:"column:is_correct"`
}

type Answer struct {
	SessionID  string     `gorm:"column:session_id"`
	QuestionID string     `gorm:"column:question_id"`
	ChoiceID   string     `gorm:"column:choice_id"`
	CreatedAt  time.Time  `gorm:"column:created_at"`
	UpdatedAt  *time.Time `gorm:"column:updated_at"`
}

type SummaryResponse struct {
	Total   int `json:"total"`
	Correct int `json:"correct"`
}
