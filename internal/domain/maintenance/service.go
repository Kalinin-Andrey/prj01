package maintenance

import (
	"context"

	"github.com/pkg/errors"

	"carizza/internal/domain"
	"carizza/internal/pkg/log"
)

const MaxLIstLimit = 1000

// IService encapsulates usecase logic for user.
type IService interface {
	NewEntity() *Maintenance
	Get(ctx context.Context, id uint) (*Maintenance, error)
	Query(ctx context.Context, query domain.DBQueryConditions) ([]Maintenance, error)
	Create(ctx context.Context, entity *Maintenance) error
	Update(ctx context.Context, entity *Maintenance) error
	Save(ctx context.Context, entity *Maintenance) error
	Delete(ctx context.Context, entity *Maintenance) error
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
func (s service) Query(ctx context.Context, query domain.DBQueryConditions) ([]Maintenance, error) {
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
func (s service) Delete(ctx context.Context, entity *Maintenance) error {
	return s.repository.Save(ctx, entity)
}
