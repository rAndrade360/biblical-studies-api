package models

import (
	"testing"
	"time"
)

func TestNewAlternative(t *testing.T) {
	type args struct {
		questionId string
		value      string
		isCorrect  bool
	}
	tests := []struct {
		name    string
		args    args
		want    *Alternative
		wantErr bool
	}{
		{
			name: "should be able to return error",
			args: args{
				questionId: "",
				value:      "A",
				isCorrect:  false,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should be able to return success",
			args: args{
				questionId: "123456789",
				value:      "B",
				isCorrect:  true,
			},
			want: &Alternative{
				QuestionID: "123456789",
				Value:      "B",
				IsCorret:   true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAlternative(tt.args.questionId, tt.args.value, tt.args.isCorrect)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAlternative() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil {
				if got == nil || got != nil && (tt.want.QuestionID != got.QuestionID || tt.want.Value != got.Value || tt.want.IsCorret != got.IsCorret) {
					t.Errorf("NewAlternative() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestAlternative_validate(t *testing.T) {
	type fields struct {
		ID         string
		QuestionID string
		Value      string
		IsCorret   bool
		CreatedAt  time.Time
		UpdatedAt  time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "should be able to return error when questionID is empty",
			fields: fields{
				Value: "123",
			},
			wantErr: true,
		},
		{
			name: "should be able to return error when value is empty",
			fields: fields{
				QuestionID: "123",
			},
			wantErr: true,
		},
		{
			name: "should be able to return success",
			fields: fields{
				QuestionID: "123",
				Value:      "123",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Alternative{
				ID:         tt.fields.ID,
				QuestionID: tt.fields.QuestionID,
				Value:      tt.fields.Value,
				IsCorret:   tt.fields.IsCorret,
				CreatedAt:  tt.fields.CreatedAt,
				UpdatedAt:  tt.fields.UpdatedAt,
			}
			if err := a.validate(); (err != nil) != tt.wantErr {
				t.Errorf("Alternative.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
