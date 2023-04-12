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
	GetById(id string) (*models.QuestionGroup, error)
	List() ([]models.QuestionGroup, error)
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

func (r *repo) GetById(id string) (*models.QuestionGroup, error) {
	row := r.db.QueryRow("SELECT qg.id, qg.name, qg.description, qg.image_url, qg.prev_qg_id, qg.created_at, qg.updated_at FROM question_groups as qg WHERE qg.id = ?;", id)

	var qg models.QuestionGroup

	err := row.Scan(
		&qg.ID,
		&qg.Name,
		&qg.Description,
		&qg.ImageUrl,
		&qg.PrevQGID,
		&qg.CreatedAt,
		&qg.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &qg, nil
}

func (r *repo) List() ([]models.QuestionGroup, error) {
	rows, err := r.db.Query("SELECT qg.id, qg.name, qg.description, qg.image_url, qg.prev_qg_id, qg.created_at, qg.updated_at FROM question_groups as qg;")
	if err != nil {
		return nil, err
	}

	var questionGroups []models.QuestionGroup
	defer rows.Close()
	for rows.Next() {
		var qg models.QuestionGroup

		err := rows.Scan(
			&qg.ID,
			&qg.Name,
			&qg.Description,
			&qg.ImageUrl,
			&qg.PrevQGID,
			&qg.CreatedAt,
			&qg.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		questionGroups = append(questionGroups, qg)
	}

	return questionGroups, nil
}
