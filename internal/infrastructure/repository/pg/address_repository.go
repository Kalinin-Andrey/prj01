package pg

import (
	"context"

	"github.com/jinzhu/gorm"

	"carizza/internal/pkg/apperror"
	minipkg_gorm "carizza/pkg/db/gorm"
	"carizza/pkg/selection_condition"

	"carizza/internal/domain/address"
)

// AddressRepository is a repository for the service entity
type AddressRepository struct {
	repository
}

var _ address.Repository = (*AddressRepository)(nil)

// NewAddressRepository creates a new c
func NewAddressRepository(repository *repository) (*AddressRepository, error) {
	r := &AddressRepository{repository: *repository}
	r.autoMigrate()
	return r, nil
}

func (r AddressRepository) autoMigrate() {
	if r.db.IsAutoMigrate() {
		r.db.DB().AutoMigrate(&address.Address{})
	}
}

// Get reads the album with the specified ID from the database.
func (r AddressRepository) Get(ctx context.Context, id uint) (*address.Address, error) {
	entity := &address.Address{}

	err := r.dbWithDefaults().First(entity, id).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return entity, apperror.ErrNotFound
		}
	}
	return entity, err
}

func (r AddressRepository) First(ctx context.Context, entity *address.Address) (*address.Address, error) {
	err := r.dbWithDefaults().Where(entity).First(entity).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return entity, apperror.ErrNotFound
		}
	}
	return entity, err
}

// Query retrieves the album records with the specified offset and limit from the database.
func (r AddressRepository) Query(ctx context.Context, cond selection_condition.SelectionCondition) ([]address.Address, error) {
	items := []address.Address{}
	db, err := minipkg_gorm.ApplyConditions(r.dbWithDefaults(), cond)
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
