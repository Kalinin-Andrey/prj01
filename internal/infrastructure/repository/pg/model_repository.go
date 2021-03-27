package pg

import (
	"context"

	"github.com/jinzhu/gorm"

	minipkg_gorm "carizza/pkg/db/gorm"
	"carizza/pkg/selection_condition"

	"carizza/internal/domain/model"
	"carizza/internal/pkg/apperror"
)

// UserRepository is a repository for the model entity
type ModelRepository struct {
	repository
}

var _ model.Repository = (*ModelRepository)(nil)

// New creates a new ModelRepository
func NewModelRepository(repository *repository) (*ModelRepository, error) {
	return &ModelRepository{repository: *repository}, nil
}

func (r ModelRepository) autoMigrate() {
	if r.db.IsAutoMigrate() {
		r.db.DB().AutoMigrate(&model.Model{})
	}
}

// Get reads the album with the specified ID from the database.
func (r ModelRepository) Get(ctx context.Context, id uint) (*model.Model, error) {
	entity := &model.Model{}

	err := r.DB().First(entity, id).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return entity, apperror.ErrNotFound
		}
	}
	return entity, err
}

func (r ModelRepository) First(ctx context.Context, entity *model.Model) (*model.Model, error) {
	err := r.DB().Where(entity).First(entity).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return entity, apperror.ErrNotFound
		}
	}
	return entity, err
}

// Query retrieves the album records with the specified offset and limit from the database.
func (r ModelRepository) Query(ctx context.Context, cond *selection_condition.SelectionCondition) ([]model.Model, error) {
	items := []model.Model{}
	db := minipkg_gorm.Conditions(r.DB().Model(&model.Model{}), cond)
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
