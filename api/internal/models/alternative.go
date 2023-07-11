package models

import (
	"time"

	"github.com/google/uuid"
)

type Alternative struct {
	ID         string    `json:"id"`
	QuestionID string    `json:"question_id"`
	Value      string    `json:"value"`
	IsCorret   bool      `json:"is_correct"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func NewAlternative(questionId, value string, isCorrect bool) Alternative {
	return Alternative{
		ID:         uuid.NewString(),
		QuestionID: questionId,
		Value:      value,
		IsCorret:   isCorrect,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}
