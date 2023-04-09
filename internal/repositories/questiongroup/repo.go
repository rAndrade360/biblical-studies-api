package questiongroup

import (
	"database/sql"

	"github.com/rAndrade360/biblical-studies-api/internal/models"
)

type repo struct {
	db *sql.DB
}

type QuestionGroupRepository interface {
	Create(qg *models.QuestionGroup) error
}

func NewQuestionGroupRepository(db *sql.DB) QuestionGroupRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(qg *models.QuestionGroup) error {
	stmt, err := r.db.Prepare("INSERT INTO question_groups VALUES (?,?,?,?,?,?,?);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(qg.ID, qg.Name, qg.Description, qg.ImageUrl, qg.PrevQGID, qg.CreatedAt, qg.UpdatedAt)
	return err
}
