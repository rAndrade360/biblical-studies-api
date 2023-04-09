package models

import "time"

type QuestionGroup struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"image_url"`
	PrevQGID    string    `json:"prev_qg_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
