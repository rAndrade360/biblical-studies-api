package models

import (
	"time"

	"github.com/google/uuid"
)

type Question struct {
	ID              string    `json:"id"`
	QuestionGroupID string    `json:"questionGroupId"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	BibleText       string    `json:"bibleText"`
	ImageUrl        string    `json:"imageUrl"`
	SortNumber      int       `json:"sortNumber"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

func NewQuestion(questionGroupID, title, description, bibleText, imageUrl string, sortNumber int) Question {
	return Question{
		ID:              uuid.NewString(),
		QuestionGroupID: questionGroupID,
		Title:           title,
		Description:     description,
		BibleText:       bibleText,
		ImageUrl:        imageUrl,
		SortNumber:      sortNumber,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}
