package maintenance

import (
	"carizza/pkg/selection_condition"
	"context"

	"github.com/pkg/errors"

	"github.com/minipkg/go-app-common/log"
)

// IService encapsulates usecase logic for user.
type IService interface {
	NewEntity() *Maintenance
	Get(ctx context.Context, id uint) (*Maintenance, error)
	Query(ctx context.Context, query *selection_condition.SelectionCondition) ([]Maintenance, error)
	Create(ctx context.Context, entity *Maintenance) error
	Update(ctx context.Context, entity *Maintenance) error
	Save(ctx context.Context, entity *Maintenance) error
	Delete(ctx context.Context, id uint) error
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

func (s service) NewEntity() *Maintenance {
	return New()
}

// Get returns the entity with the specified ID.
func (s service) Get(ctx context.Context, id uint) (*Maintenance, error) {
	entity, err := s.repository.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

// Query returns the items with the specified offset and limit.
func (s service) Query(ctx context.Context, query *selection_condition.SelectionCondition) ([]Maintenance, error) {
	items, err := s.repository.Query(ctx, query)
	if err != nil {
		return nil, errors.Wrapf(err, "Can not find a list of Maintenance by query: %v", query)
	}
	return items, nil
}

// Create saves a new Maintenance record in the database.
func (s service) Create(ctx context.Context, entity *Maintenance) error {
	return s.repository.Create(ctx, entity)
}

// Update saves an existing Maintenance record in the database.
func (s service) Update(ctx context.Context, entity *Maintenance) error {
	return s.repository.Update(ctx, entity)
}

// Save saves a Maintenance record in the database.
func (s service) Save(ctx context.Context, entity *Maintenance) error {
	return s.repository.Save(ctx, entity)
}

// Delete (soft) deletes a Maintenance record in the database.
func (s service) Delete(ctx context.Context, id uint) error {
	return s.repository.Delete(ctx, id)
}
