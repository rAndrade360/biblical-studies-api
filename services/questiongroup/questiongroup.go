package questiongroup

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/rAndrade360/biblical-studies-api/internal/models"
	repositories "github.com/rAndrade360/biblical-studies-api/internal/repositories/questiongroup"
	errors "github.com/rAndrade360/biblical-studies-api/pkg/error"
	"github.com/rAndrade360/biblical-studies-api/pkg/logger"
)

type service struct {
	repo repositories.QuestionGroupRepository
	log  logger.Logger
}

type QuestionGroupService interface {
	Create(qg *models.QuestionGroup) error
	GetById(id string) (*models.QuestionGroup, error)
	List() ([]models.QuestionGroup, error)
}

func NewQuestionGroupService(repo repositories.QuestionGroupRepository, logger logger.Logger) QuestionGroupService {
	return &service{
		repo: repo,
		log:  logger,
	}
}

func (s *service) Create(qg *models.QuestionGroup) error {
	jb, _ := json.Marshal(qg)
	s.log.Infof("Creating QuestionGroup with the data: %s", string(jb))

	qg.ID = uuid.New().String()
	qg.CreatedAt = time.Now()
	qg.UpdatedAt = time.Now()

	return s.repo.Create(qg)
}

func (s *service) GetById(id string) (*models.QuestionGroup, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		s.log.Errorf("Error to get questionGroup: %s", err.Error())
		return nil, errors.INVALIDINPUT
	}

	s.log.Infof("Getting QuestionGroup by id: %s", id)

	return s.repo.GetById(id)
}

func (s *service) List() ([]models.QuestionGroup, error) {
	s.log.Infof("List QuestionGroups")

	return s.repo.List()
}
