package order

import (
	"context"
	"github.com/minipkg/selection_condition"

	"github.com/pkg/errors"

	"github.com/minipkg/log"
)

// IService encapsulates usecase logic for user.
type IService interface {
	NewEntity() *Order
	Get(ctx context.Context, id uint) (*Order, error)
	Query(ctx context.Context, query *selection_condition.SelectionCondition) ([]Order, error)
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
func (s service) defaultConditions() *selection_condition.SelectionCondition {
	return &selection_condition.SelectionCondition{}
}

func (s service) NewEntity() *Order {
	return New()
}

// Get returns the entity with the specified ID.
func (s service) Get(ctx context.Context, id uint) (*Order, error) {
	entity, err := s.repository.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

// Query returns the items with the specified offset and limit.
func (s service) Query(ctx context.Context, query *selection_condition.SelectionCondition) ([]Order, error) {
	items, err := s.repository.Query(ctx, query)
	if err != nil {
		return nil, errors.Wrapf(err, "Can not find a list of Order by query: %v", query)
	}
	return items, nil
}
