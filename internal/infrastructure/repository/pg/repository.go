package pg

import (
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
	minipkg_gorm "github.com/minipkg/db/gorm"
	"github.com/minipkg/selection_condition"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"github.com/minipkg/log"
)

// IRepository is an interface of repository
type IRepository interface{}

// repository persists albums in database
type repository struct {
	db         minipkg_gorm.IDB
	logger     log.ILogger
	Conditions *selection_condition.SelectionCondition
}

const DefaultLimit = 1000

// GetRepository return a repository
func GetRepository(logger log.ILogger, dbase minipkg_gorm.IDB, entity string) (repo IRepository, err error) {
	r := &repository{
		db:     dbase,
		logger: logger,
	}

	switch entity {
	case user.EntityName:
		repo, err = NewUserRepository(r)
	case mark.EntityName:
		repo, err = NewMarkRepository(r)
	case model.EntityName:
		repo, err = NewModelRepository(r)
	case generation.EntityName:
		repo, err = NewGenerationRepository(r)
	case serie.EntityName:
		repo, err = NewSerieRepository(r)
	case modification.EntityName:
		repo, err = NewModificationRepository(r)

	case maintenance.EntityName:
		repo, err = NewMaintenanceRepository(r)
	case work.EntityName:
		repo, err = NewWorkRepository(r)
	case supply.EntityName:
		repo, err = NewSupplyRepository(r)
	case order.EntityName:
		repo, err = NewOrderRepository(r)
	case car.EntityName:
		repo, err = NewCarRepository(r)
	case address.EntityName:
		repo, err = NewAddressRepository(r)
	case client.EntityName:
		repo, err = NewClientRepository(r)
	default:
		err = errors.Errorf("Repository for entity %q not found", entity)
	}
	return repo, err
}

func (r *repository) SetDefaultConditions(defaultConditions *selection_condition.SelectionCondition) {
	r.Conditions = defaultConditions
}

func (r repository) DB() *gorm.DB {
	return minipkg_gorm.Conditions(r.db.DB(), r.Conditions)
}
