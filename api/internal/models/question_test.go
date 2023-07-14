package models

import (
	"reflect"
	"testing"
	"time"
)

func TestNewQuestion(t *testing.T) {
	type args struct {
		questionGroupID string
		title           string
		description     string
		bibleText       string
		imageUrl        string
		sortNumber      int
	}
	tests := []struct {
		name    string
		args    args
		want    *Question
		wantErr bool
	}{
		{
			name: "should be able to return error when validation fails",
			args: args{
				questionGroupID: "",
				title:           "Bla",
				description:     "",
				bibleText:       "",
				imageUrl:        "",
				sortNumber:      1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "should be able to return a valida question",
			args: args{
				questionGroupID: "123456",
				title:           "Bla",
				description:     "blr",
				bibleText:       "ble",
				imageUrl:        "",
				sortNumber:      10,
			},
			want: &Question{
				ID:              "",
				QuestionGroupID: "123456",
				Title:           "Bla",
				Description:     "blr",
				BibleText:       "ble",
				ImageUrl:        "",
				SortNumber:      10,
				CreatedAt:       time.Time{},
				UpdatedAt:       time.Time{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewQuestion(tt.args.questionGroupID, tt.args.title, tt.args.description, tt.args.bibleText, tt.args.imageUrl, tt.args.sortNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewQuestion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil && tt.want != nil {
				tt.want.ID = got.ID
				tt.want.CreatedAt = got.CreatedAt
				tt.want.UpdatedAt = got.UpdatedAt
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQuestion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuestion_validate(t *testing.T) {
	type fields struct {
		ID              string
		QuestionGroupID string
		Title           string
		Description     string
		BibleText       string
		ImageUrl        string
		SortNumber      int
		CreatedAt       time.Time
		UpdatedAt       time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "should be able to return error when QuestionGroupID is empty",
			fields: fields{
				ID:              "",
				QuestionGroupID: "",
				Title:           "A terra prometida",
				Description:     "",
				BibleText:       "",
				ImageUrl:        "",
				SortNumber:      1,
				CreatedAt:       time.Time{},
				UpdatedAt:       time.Time{},
			},
			wantErr: true,
		},
		{
			name: "should be able to return error when Title is empty",
			fields: fields{
				ID:              "",
				QuestionGroupID: "123456",
				Title:           "",
				Description:     "",
				BibleText:       "",
				ImageUrl:        "",
				SortNumber:      1,
				CreatedAt:       time.Time{},
				UpdatedAt:       time.Time{},
			},
			wantErr: true,
		},
		{
			name: "should be able to return error when SortNumber is lt 1",
			fields: fields{
				ID:              "",
				QuestionGroupID: "123456",
				Title:           "A terra prometida",
				Description:     "",
				BibleText:       "",
				ImageUrl:        "",
				SortNumber:      -1,
				CreatedAt:       time.Time{},
				UpdatedAt:       time.Time{},
			},
			wantErr: true,
		},
		{
			name: "should be able to return error == nil",
			fields: fields{
				ID:              "",
				QuestionGroupID: "123456",
				Title:           "A terra prometida",
				Description:     "",
				BibleText:       "",
				ImageUrl:        "",
				SortNumber:      3,
				CreatedAt:       time.Time{},
				UpdatedAt:       time.Time{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Question{
				ID:              tt.fields.ID,
				QuestionGroupID: tt.fields.QuestionGroupID,
				Title:           tt.fields.Title,
				Description:     tt.fields.Description,
				BibleText:       tt.fields.BibleText,
				ImageUrl:        tt.fields.ImageUrl,
				SortNumber:      tt.fields.SortNumber,
				CreatedAt:       tt.fields.CreatedAt,
				UpdatedAt:       tt.fields.UpdatedAt,
			}
			if err := q.validate(); (err != nil) != tt.wantErr {
				t.Errorf("Question.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
