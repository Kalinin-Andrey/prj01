package work

import (
	"context"

	"github.com/pkg/errors"

	"carizza/internal/domain"
	"carizza/internal/pkg/log"
)

const MaxLIstLimit = 1000

// IService encapsulates usecase logic for user.
type IService interface {
	NewEntity() *Work
	Get(ctx context.Context, id uint) (*Work, error)
	Query(ctx context.Context, query domain.DBQueryConditions) ([]Work, error)
}

type service struct {
	logger     log.ILogger
	repository Repository
}

// NewService creates a new service.
func NewService(logger log.ILogger, repo Repository) IService {
	s := &service{
		logger:     logger,
		repository: repo,
	}
	repo.SetDefaultConditions(s.defaultConditions())
	return s
}

// Defaults returns defaults params
func (s service) defaultConditions() domain.DBQueryConditions {
	return domain.DBQueryConditions{}
}

func (s service) NewEntity() *Work {
	return New()
}

// Get returns the entity with the specified ID.
func (s service) Get(ctx context.Context, id uint) (*Work, error) {
	entity, err := s.repository.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

// Query returns the items with the specified offset and limit.
func (s service) Query(ctx context.Context, query domain.DBQueryConditions) ([]Work, error) {
	items, err := s.repository.Query(ctx, query)
	if err != nil {
		return nil, errors.Wrapf(err, "Can not find a list of Work by query: %v", query)
	}
	return items, nil
}
