package questiongroup

import (
	"context"

	"github.com/google/uuid"
	"github.com/rAndrade360/biblical-studies-api/api/internal/models"
	repositories "github.com/rAndrade360/biblical-studies-api/api/internal/repositories/questiongroup"
	errors "github.com/rAndrade360/biblical-studies-api/pkg/error"
	"github.com/rAndrade360/biblical-studies-api/pkg/logger"
)

type service struct {
	repo repositories.QuestionGroupRepository
}

type QuestionGroupService interface {
	Create(ctx context.Context, qg *models.QuestionGroup) error
	GetById(ctx context.Context, id string) (*models.QuestionGroup, error)
	List(ctx context.Context) ([]models.QuestionGroup, error)
}

func NewQuestionGroupService(repo repositories.QuestionGroupRepository) QuestionGroupService {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(ctx context.Context, qg *models.QuestionGroup) error {
	log := logger.GetLoggerCtx(ctx)

	log.Info("Creating QuestionGroup", qg)

	return s.repo.Create(qg)
}

func (s *service) GetById(ctx context.Context, id string) (*models.QuestionGroup, error) {
	log := logger.GetLoggerCtx(ctx)

	_, err := uuid.Parse(id)
	if err != nil {
		log.Error("Error to get questionGroup", err.Error())
		return nil, errors.INVALIDINPUT
	}

	log.Info("Getting QuestionGroup by id", id)

	return s.repo.GetById(id)
}

func (s *service) List(ctx context.Context) ([]models.QuestionGroup, error) {
	log := logger.GetLoggerCtx(ctx)

	log.Info("List QuestionGroups")

	return s.repo.List()
}
