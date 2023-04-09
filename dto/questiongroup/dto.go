package questiongroup

import "time"

type QuestionGroupHttpCreate struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ImageUrl    string `json:"image_url,omitempty"`
	PrevQGID    string `json:"prev_qg_id,omitempty"`
}

type QuestionGroupHttpCreateResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"image_url"`
	PrevQGID    string    `json:"prev_qg_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
