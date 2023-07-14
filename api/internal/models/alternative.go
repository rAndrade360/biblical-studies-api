package models

import (
	"errors"
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

func NewAlternative(questionId, value string, isCorrect bool) (*Alternative, error) {
	a := Alternative{
		ID:         uuid.NewString(),
		QuestionID: questionId,
		Value:      value,
		IsCorret:   isCorrect,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	err := a.validate()
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (a *Alternative) validate() error {
	if len(a.QuestionID) == 0 || len(a.Value) == 0 {
		return errors.New("invalid params")
	}

	return nil
}
