package pg

import (
	"context"

	"github.com/jinzhu/gorm"

	"carizza/internal/pkg/apperror"

	"carizza/internal/domain"
	"carizza/internal/domain/supply"
)

// SupplyRepository is a repository for the service entity
type SupplyRepository struct {
	repository
}

var _ supply.Repository = (*SupplyRepository)(nil)

// NewSupplyRepository creates a new c
func NewSupplyRepository(repository *repository) (*SupplyRepository, error) {
	r := &SupplyRepository{repository: *repository}
	r.autoMigrate()
	return r, nil
}

func (r SupplyRepository) autoMigrate() {
	if r.db.IsAutoMigrate() {
		r.db.DB().AutoMigrate(&supply.Supply{})
	}
}

// Get reads the album with the specified ID from the database.
func (r SupplyRepository) Get(ctx context.Context, id uint) (*supply.Supply, error) {
	entity := &supply.Supply{}

	err := r.dbWithDefaults().First(entity, id).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return entity, apperror.ErrNotFound
		}
	}
	return entity, err
}

func (r SupplyRepository) First(ctx context.Context, entity *supply.Supply) (*supply.Supply, error) {
	err := r.dbWithDefaults().Where(entity).First(entity).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return entity, apperror.ErrNotFound
		}
	}
	return entity, err
}

// Query retrieves the album records with the specified offset and limit from the database.
func (r SupplyRepository) Query(ctx context.Context, cond domain.DBQueryConditions) ([]supply.Supply, error) {
	items := []supply.Supply{}
	db, err := r.applyConditions(r.dbWithDefaults(), cond)
	if err != nil {
		return nil, err
	}

	err = db.Find(&items).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return items, apperror.ErrNotFound
		}
	}
	return items, err
}
