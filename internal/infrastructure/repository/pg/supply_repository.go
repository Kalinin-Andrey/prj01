package pg

import (
	"context"

	"github.com/jinzhu/gorm"

	minipkg_gorm "github.com/minipkg/db/gorm"
	"github.com/minipkg/selection_condition"

	"carizza/internal/pkg/apperror"

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

	err := r.DB().First(entity, id).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return entity, apperror.ErrNotFound
		}
	}
	return entity, err
}

func (r SupplyRepository) First(ctx context.Context, entity *supply.Supply) (*supply.Supply, error) {
	err := r.DB().Where(entity).First(entity).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return entity, apperror.ErrNotFound
		}
	}
	return entity, err
}

// Query retrieves the album records with the specified offset and limit from the database.
func (r SupplyRepository) Query(ctx context.Context, cond *selection_condition.SelectionCondition) ([]supply.Supply, error) {
	items := []supply.Supply{}
	db := minipkg_gorm.Conditions(r.DB(), cond)
	if db.Error != nil {
		return nil, db.Error
	}

	err := db.Find(&items).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return items, apperror.ErrNotFound
		}
	}
	return items, err
}
