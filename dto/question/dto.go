package question

import "time"

type QuestionHttpCreate struct {
	QuestionGroupID string `json:"questionGroupId,omitempty"`
	Title           string `json:"title,omitempty"`
	BibleText       string `json:"bibleText,omitempty"`
	Description     string `json:"description,omitempty"`
	ImageUrl        string `json:"imageUrl,omitempty"`
	SortNumber      int    `json:"sortNumber,omitempty"`
}

type QuestionHttpCreateResponse struct {
	ID              string    `json:"id"`
	QuestionGroupID string    `json:"questionGroupId,omitempty"`
	Title           string    `json:"title,omitempty"`
	Description     string    `json:"description,omitempty"`
	BibleText       string    `json:"bibleText,omitempty"`
	ImageUrl        string    `json:"imageUrl,omitempty"`
	SortNumber      int       `json:"sortNumber,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
