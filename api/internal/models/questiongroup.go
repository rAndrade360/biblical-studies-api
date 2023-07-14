package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type QuestionGroup struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"imageUrl"`
	SortNumber  int       `json:"sortNumber"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewQuestionGroup(name, description, imageUrl string, sortNumber int) (*QuestionGroup, error) {
	qg := QuestionGroup{
		ID:          uuid.NewString(),
		Name:        name,
		Description: description,
		ImageUrl:    imageUrl,
		SortNumber:  sortNumber,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := qg.validate()
	if err != nil {
		return nil, err
	}

	return &qg, nil
}

func (qg *QuestionGroup) validate() error {
	if len(qg.Name) == 0 || qg.SortNumber <= 0 {
		return errors.New("invalid params")
	}

	return nil
}
