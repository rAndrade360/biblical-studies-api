package alternative

import (
	"database/sql"

	"github.com/rAndrade360/biblical-studies-api/internal/models"
)

type repo struct {
	db *sql.DB
}

type AlternativeRepository interface {
	Create(q *models.Alternative) error
	GetById(id string) (*models.Alternative, error)
	List() ([]models.Alternative, error)
	GetByQuestionId(id string) ([]models.Alternative, error)
}

func NewAlternativeRepository(db *sql.DB) AlternativeRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(a *models.Alternative) error {
	stmt, err := r.db.Prepare("INSERT INTO alternatives VALUES (?,?,?,?,?,?);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		a.ID,
		a.QuestionID,
		a.Value,
		a.IsCorret,
		a.CreatedAt,
		a.UpdatedAt,
	)

	return err
}

func (r *repo) GetById(id string) (*models.Alternative, error) {
	row := r.db.QueryRow(`SELECT
	a.id,
	a.question_id,
	a.value,
	a.is_correct,
	a.created_at,
	a.updated_at FROM alternatives as a WHERE a.id = ?;`, id)

	var a models.Alternative

	err := row.Scan(
		&a.ID,
		&a.QuestionID,
		&a.Value,
		&a.IsCorret,
		&a.CreatedAt,
		&a.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (r *repo) GetByQuestionId(id string) ([]models.Alternative, error) {
	rows, err := r.db.Query(`SELECT
	a.id,
	a.question_id,
	a.value,
	a.is_correct,
	a.created_at,
	a.updated_at FROM alternatives as a WHERE a.question_id = ?;`, id)
	if err != nil {
		return nil, err
	}

	var alternatives []models.Alternative
	defer rows.Close()
	for rows.Next() {
		var a models.Alternative

		err := rows.Scan(
			&a.ID,
			&a.QuestionID,
			&a.Value,
			&a.IsCorret,
			&a.CreatedAt,
			&a.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		alternatives = append(alternatives, a)
	}

	return alternatives, nil
}

func (r *repo) List() ([]models.Alternative, error) {
	rows, err := r.db.Query(`SELECT
	a.id,
	a.question_id,
	a.value,
	a.is_correct,
	a.created_at,
	a.updated_at FROM alternatives as a;`)
	if err != nil {
		return nil, err
	}

	var alternatives []models.Alternative
	defer rows.Close()
	for rows.Next() {
		var a models.Alternative

		err := rows.Scan(
			&a.ID,
			&a.QuestionID,
			&a.Value,
			&a.IsCorret,
			&a.CreatedAt,
			&a.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		alternatives = append(alternatives, a)
	}

	return alternatives, nil
}
