package pg

import (
	"context"

	"github.com/jinzhu/gorm"

	"carizza/internal/pkg/apperror"

	"carizza/internal/domain"
	"carizza/internal/domain/car"
)

// CarRepository is a repository for the service entity
type CarRepository struct {
	repository
}

var _ car.Repository = (*CarRepository)(nil)

// NewCarRepository creates a new c
func NewCarRepository(repository *repository) (*CarRepository, error) {
	r := &CarRepository{repository: *repository}
	r.autoMigrate()
	return r, nil
}

func (r CarRepository) autoMigrate() {
	if r.db.IsAutoMigrate() {
		r.db.DB().AutoMigrate(&car.Car{})
	}
}

// Get reads the album with the specified ID from the database.
func (r CarRepository) Get(ctx context.Context, id uint) (*car.Car, error) {
	entity := &car.Car{}

	err := r.dbWithDefaults().First(entity, id).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return entity, apperror.ErrNotFound
		}
	}
	return entity, err
}

func (r CarRepository) First(ctx context.Context, entity *car.Car) (*car.Car, error) {
	err := r.dbWithDefaults().Where(entity).First(entity).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return entity, apperror.ErrNotFound
		}
	}
	return entity, err
}

// Query retrieves the album records with the specified offset and limit from the database.
func (r CarRepository) Query(ctx context.Context, cond domain.DBQueryConditions) ([]car.Car, error) {
	items := []car.Car{}
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