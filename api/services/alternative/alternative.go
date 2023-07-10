package alternative

import (
	"context"

	"github.com/google/uuid"
	"github.com/rAndrade360/biblical-studies-api/api/internal/models"
	repositories "github.com/rAndrade360/biblical-studies-api/api/internal/repositories/alternative"
	questionservice "github.com/rAndrade360/biblical-studies-api/api/services/question"
	errors "github.com/rAndrade360/biblical-studies-api/pkg/error"
	"github.com/rAndrade360/biblical-studies-api/pkg/logger"
)

type service struct {
	repo            repositories.AlternativeRepository
	questionService questionservice.QuestionService
}

type AlternativeService interface {
	Create(ctx context.Context, q *models.Alternative) error
	GetById(ctx context.Context, id string) (*models.Alternative, error)
	List(ctx context.Context) ([]models.Alternative, error)
	GetByQuestionId(ctx context.Context, id string) ([]models.Alternative, error)
}

func NewAlternativeService(repo repositories.AlternativeRepository, questionService questionservice.QuestionService) AlternativeService {
	return &service{
		repo:            repo,
		questionService: questionService,
	}
}

func (s *service) Create(ctx context.Context, a *models.Alternative) error {
	log := logger.GetLoggerCtx(ctx)

	_, err := s.questionService.GetById(ctx, a.QuestionID)
	if err != nil {
		return err
	}

	log.Info("Creating Alternative", a)

	return s.repo.Create(a)
}

func (s *service) GetById(ctx context.Context, id string) (*models.Alternative, error) {
	log := logger.GetLoggerCtx(ctx)

	_, err := uuid.Parse(id)
	if err != nil {
		log.Error("Error to get Alternative", err.Error())
		return nil, errors.INVALIDINPUT
	}

	log.Info("Getting Alternative by id", id)

	return s.repo.GetById(id)
}

func (s *service) GetByQuestionId(ctx context.Context, id string) ([]models.Alternative, error) {
	log := logger.GetLoggerCtx(ctx)

	_, err := uuid.Parse(id)
	if err != nil {
		log.Error("Error to get Alternative", err.Error())
		return nil, errors.INVALIDINPUT
	}

	log.Info("Getting Alternative by question id", id)

	return s.repo.GetByQuestionId(id)
}

func (s *service) List(ctx context.Context) ([]models.Alternative, error) {
	log := logger.GetLoggerCtx(ctx)

	log.Info("List Alternatives")

	return s.repo.List()
}
