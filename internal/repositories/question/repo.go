package question

import (
	"database/sql"

	"github.com/rAndrade360/biblical-studies-api/internal/models"
)

type repo struct {
	db *sql.DB
}

type QuestionRepository interface {
	Create(q *models.Question) error
	GetById(id string) (*models.Question, error)
	List() ([]models.Question, error)
}

func NewQuestionRepository(db *sql.DB) QuestionRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(q *models.Question) error {
	stmt, err := r.db.Prepare("INSERT INTO questions VALUES (?,?,?,?,?,?,?,?,?);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		q.ID,
		q.QuestionGroupID,
		q.Title,
		q.Description,
		q.BibleText,
		q.ImageUrl,
		q.SortNumber,
		q.CreatedAt,
		q.UpdatedAt,
	)

	return err
}

func (r *repo) GetById(id string) (*models.Question, error) {
	row := r.db.QueryRow(`SELECT
	q.id,
	q.question_group_id,
	q.title,
	q.description,
	q.bible_text,
	q.image_url,
	q.sort_number,
	q.created_at,
	q.updated_at FROM questions as q WHERE q.id = ?;`, id)

	var q models.Question

	err := row.Scan(
		&q.ID,
		&q.QuestionGroupID,
		&q.Title,
		&q.Description,
		&q.BibleText,
		&q.ImageUrl,
		&q.SortNumber,
		&q.CreatedAt,
		&q.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &q, nil
}

func (r *repo) List() ([]models.Question, error) {
	rows, err := r.db.Query(`SELECT
	q.id,
	q.question_group_id,
	q.title,
	q.description,
	q.bible_text,
	q.image_url,
	q.sort_number,
	q.created_at,
	q.updated_at FROM questions as q;`)
	if err != nil {
		return nil, err
	}

	var questions []models.Question
	defer rows.Close()
	for rows.Next() {
		var q models.Question

		err := rows.Scan(
			&q.ID,
			&q.QuestionGroupID,
			&q.Title,
			&q.Description,
			&q.BibleText,
			&q.ImageUrl,
			&q.SortNumber,
			&q.CreatedAt,
			&q.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		questions = append(questions, q)
	}

	return questions, nil
}
