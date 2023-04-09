package questiongroup

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rAndrade360/biblical-studies-api/internal/models"
	repositories "github.com/rAndrade360/biblical-studies-api/internal/repositories/questiongroup"
)

type service struct {
	repo repositories.QuestionGroupRepository
}

type QuestionGroupService interface {
	Create(qg *models.QuestionGroup) error
}

func NewQuestionGroupService(repo repositories.QuestionGroupRepository) QuestionGroupService {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(qg *models.QuestionGroup) error {
	fmt.Printf("Creating QuestionGroup with the data: %#v\n", qg)
	qg.ID = uuid.New().String()
	qg.CreatedAt = time.Now()
	qg.UpdatedAt = time.Now()

	return s.repo.Create(qg)
}
