package gorm

import (
	"context"

	"github.com/pkg/errors"

	minipkg_gorm "github.com/minipkg/db/gorm"
	"github.com/minipkg/log"
	"github.com/minipkg/selection_condition"
	"gorm.io/gorm"

	"carizza/internal/domain/address"
	"carizza/internal/domain/car"
	"carizza/internal/domain/client"
	"carizza/internal/domain/generation"
	"carizza/internal/domain/maintenance"
	"carizza/internal/domain/mark"
	"carizza/internal/domain/model"
	"carizza/internal/domain/modification"
	"carizza/internal/domain/order"
	"carizza/internal/domain/serie"
	"carizza/internal/domain/supply"
	"carizza/internal/domain/user"
	"carizza/internal/domain/work"
)

// IRepository is an interface of repository
type IRepository interface {
	DB() *gorm.DB
}

// repository persists albums in database
type repository struct {
	db         minipkg_gorm.IDB
	logger     log.ILogger
	Conditions *selection_condition.SelectionCondition
	model      interface{}
}

// GetRepository return a repository
func GetRepository(logger log.ILogger, dbase minipkg_gorm.IDB, entity string) (repo IRepository, err error) {
	r := &repository{
		logger: logger,
	}

	switch entity {
	case user.EntityName:
		r.model = user.New()

		if r.db, err = dbase.WithContext(context.Background()).Model(r.model); err != nil {
			return nil, err
		}
		repo, err = NewUserRepository(r)
	case mark.EntityName:
		r.model = mark.New()

		if r.db, err = dbase.WithContext(context.Background()).Model(r.model); err != nil {
			return nil, err
		}
		repo, err = NewMarkRepository(r)
	case model.EntityName:
		r.model = model.New()

		if r.db, err = dbase.WithContext(context.Background()).Model(r.model); err != nil {
			return nil, err
		}
		repo, err = NewModelRepository(r)
	case generation.EntityName:
		r.model = generation.New()

		if r.db, err = dbase.WithContext(context.Background()).Model(r.model); err != nil {
			return nil, err
		}
		repo, err = NewGenerationRepository(r)
	case serie.EntityName:
		r.model = serie.New()

		if r.db, err = dbase.WithContext(context.Background()).Model(r.model); err != nil {
			return nil, err
		}
		repo, err = NewSerieRepository(r)
	case modification.EntityName:
		r.model = modification.New()

		if r.db, err = dbase.WithContext(context.Background()).Model(r.model); err != nil {
			return nil, err
		}
		repo, err = NewModificationRepository(r)

	case maintenance.EntityName:
		r.model = maintenance.New()

		if r.db, err = dbase.WithContext(context.Background()).Model(r.model); err != nil {
			return nil, err
		}
		repo, err = NewMaintenanceRepository(r)
	case work.EntityName:
		r.model = work.New()

		if r.db, err = dbase.WithContext(context.Background()).Model(r.model); err != nil {
			return nil, err
		}
		repo, err = NewWorkRepository(r)
	case supply.EntityName:
		r.model = supply.New()

		if r.db, err = dbase.WithContext(context.Background()).Model(r.model); err != nil {
			return nil, err
		}
		repo, err = NewSupplyRepository(r)
	case order.EntityName:
		r.model = order.New()

		if r.db, err = dbase.WithContext(context.Background()).Model(r.model); err != nil {
			return nil, err
		}
		repo, err = NewOrderRepository(r)
	case car.EntityName:
		r.model = car.New()

		if r.db, err = dbase.WithContext(context.Background()).Model(r.model); err != nil {
			return nil, err
		}
		repo, err = NewCarRepository(r)
	case address.EntityName:
		r.model = address.New()

		if r.db, err = dbase.WithContext(context.Background()).Model(r.model); err != nil {
			return nil, err
		}
		repo, err = NewAddressRepository(r)
	case client.EntityName:
		r.model = client.New()

		if r.db, err = dbase.WithContext(context.Background()).Model(r.model); err != nil {
			return nil, err
		}
		repo, err = NewClientRepository(r)
	default:
		err = errors.Errorf("Repository for entity %q not found", entity)
	}
	return repo, err
}

func (r *repository) SetDefaultConditions(defaultConditions *selection_condition.SelectionCondition) {
	r.Conditions = defaultConditions
}

func (r *repository) DB() *gorm.DB {
	return minipkg_gorm.Conditions(r.db.DB(), r.Conditions)
}
