package pg

import (
	"context"

	"github.com/jinzhu/gorm"

	"carizza/internal/pkg/apperror"

	"carizza/internal/domain"
	"carizza/internal/domain/modification"
)

// ModificationRepository is a repository for the modification entity
type ModificationRepository struct {
	repository
}

var _ modification.Repository = (*ModificationRepository)(nil)

// NewModificationRepository creates a new ModificationRepository
func NewModificationRepository(repository *repository) (*ModificationRepository, error) {
	return &ModificationRepository{repository: *repository}, nil
}

func (r ModificationRepository) autoMigrate() {
	if r.db.IsAutoMigrate() {
		r.db.DB().AutoMigrate(&modification.Modification{})
	}
}

// Get reads the album with the specified ID from the database.
func (r ModificationRepository) Get(ctx context.Context, id uint) (*modification.Modification, error) {
	entity := &modification.Modification{}

	err := r.dbWithDefaults().First(entity, id).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return entity, apperror.ErrNotFound
		}
	}
	return entity, err
}

func (r ModificationRepository) First(ctx context.Context, entity *modification.Modification) (*modification.Modification, error) {
	err := r.dbWithDefaults().Where(entity).First(entity).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return entity, apperror.ErrNotFound
		}
	}
	return entity, err
}

// Query retrieves the album records with the specified offset and limit from the database.
func (r ModificationRepository) Query(ctx context.Context, cond domain.DBQueryConditions) ([]modification.Modification, error) {
	items := []modification.Modification{}
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