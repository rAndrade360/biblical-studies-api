package models

import (
	"reflect"
	"testing"
	"time"
)

func TestNewQuestionGroup(t *testing.T) {
	type args struct {
		name        string
		description string
		imageUrl    string
		sortNumber  int
	}
	tests := []struct {
		name    string
		args    args
		want    *QuestionGroup
		wantErr bool
	}{
		{
			name: "should be able to return error when validation fails",
			args: args{
				name:        "A teoria dos teoricos",
				description: "Toda teoria eh teoricamete teorica",
				imageUrl:    "bla",
				sortNumber:  0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "",
			args: args{
				name:        "A teoria dos teoricos",
				description: "Toda teoria eh teoricamete teorica",
				imageUrl:    "bla",
				sortNumber:  1,
			},
			want: &QuestionGroup{
				ID:          "",
				Name:        "A teoria dos teoricos",
				Description: "Toda teoria eh teoricamete teorica",
				ImageUrl:    "bla",
				SortNumber:  1,
				CreatedAt:   time.Time{},
				UpdatedAt:   time.Time{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewQuestionGroup(tt.args.name, tt.args.description, tt.args.imageUrl, tt.args.sortNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewQuestionGroup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil && tt.want != nil {
				tt.want.ID = got.ID
				tt.want.CreatedAt = got.CreatedAt
				tt.want.UpdatedAt = got.UpdatedAt
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQuestionGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuestionGroup_validate(t *testing.T) {
	type fields struct {
		ID          string
		Name        string
		Description string
		ImageUrl    string
		SortNumber  int
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "should be able to return error when Name is empty",
			fields: fields{
				ID:          "",
				Name:        "",
				Description: "",
				ImageUrl:    "",
				SortNumber:  2,
				CreatedAt:   time.Time{},
				UpdatedAt:   time.Time{},
			},
			wantErr: true,
		},
		{
			name: "should be able to return error when SortNumber is less or equal to 0",
			fields: fields{
				ID:          "",
				Name:        "bla",
				Description: "",
				ImageUrl:    "",
				SortNumber:  -1,
				CreatedAt:   time.Time{},
				UpdatedAt:   time.Time{},
			},
			wantErr: true,
		},
		{
			name: "should be able to return error == nil",
			fields: fields{
				ID:          "",
				Name:        "bla",
				Description: "",
				ImageUrl:    "",
				SortNumber:  1,
				CreatedAt:   time.Time{},
				UpdatedAt:   time.Time{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qg := &QuestionGroup{
				ID:          tt.fields.ID,
				Name:        tt.fields.Name,
				Description: tt.fields.Description,
				ImageUrl:    tt.fields.ImageUrl,
				SortNumber:  tt.fields.SortNumber,
				CreatedAt:   tt.fields.CreatedAt,
				UpdatedAt:   tt.fields.UpdatedAt,
			}
			if err := qg.validate(); (err != nil) != tt.wantErr {
				t.Errorf("QuestionGroup.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
