package question

import (
	"context"

	"github.com/google/uuid"
	"github.com/rAndrade360/biblical-studies-api/internal/models"
	repositories "github.com/rAndrade360/biblical-studies-api/internal/repositories/question"
	errors "github.com/rAndrade360/biblical-studies-api/pkg/error"
	"github.com/rAndrade360/biblical-studies-api/pkg/logger"
	questiongroupservice "github.com/rAndrade360/biblical-studies-api/services/questiongroup"
)

type service struct {
	repo                 repositories.QuestionRepository
	questionGroupService questiongroupservice.QuestionGroupService
}

type QuestionService interface {
	Create(ctx context.Context, q *models.Question) error
	GetById(ctx context.Context, id string) (*models.Question, error)
	List(ctx context.Context) ([]models.Question, error)
}

func NewQuestionService(repo repositories.QuestionRepository, questionGroupService questiongroupservice.QuestionGroupService) QuestionService {
	return &service{
		repo:                 repo,
		questionGroupService: questionGroupService,
	}
}

func (s *service) Create(ctx context.Context, q *models.Question) error {
	log := logger.GetLoggerCtx(ctx)

	_, err := s.questionGroupService.GetById(ctx, q.QuestionGroupID)
	if err != nil {
		return err
	}

	log.Info("Creating Question", q)

	return s.repo.Create(q)
}

func (s *service) GetById(ctx context.Context, id string) (*models.Question, error) {
	log := logger.GetLoggerCtx(ctx)

	_, err := uuid.Parse(id)
	if err != nil {
		log.Error("Error to get question", err.Error())
		return nil, errors.INVALIDINPUT
	}

	log.Info("Getting Question by id", id)

	return s.repo.GetById(id)
}

func (s *service) List(ctx context.Context) ([]models.Question, error) {
	log := logger.GetLoggerCtx(ctx)

	log.Info("List Questions")

	return s.repo.List()
}
