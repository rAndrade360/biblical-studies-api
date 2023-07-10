package alternative

type AlternativeHttpRequest struct {
	Value    string `json:"value,omitempty"`
	IsCorret bool   `json:"is_correct,omitempty"`
}

type AlternativeHttpResponse struct {
	ID         string `json:"id,omitempty"`
	Value      string `json:"value,omitempty"`
	QuestionID string `json:"question_id,omitempty"`
}
