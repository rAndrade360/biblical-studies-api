package questiongroup

import "time"

type QuestionGroupHttpCreate struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ImageUrl    string `json:"imageUrl,omitempty"`
	SortNumber  int    `json:"sortNumber,omitempty"`
}

type QuestionGroupHttpCreateResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"imageUrl"`
	SortNumber  int       `json:"sortNumber"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
