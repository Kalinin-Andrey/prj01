package pg

import (
	"context"

	"github.com/jinzhu/gorm"

	minipkg_gorm "carizza/pkg/db/gorm"
	"carizza/pkg/selection_condition"

	"carizza/internal/domain/mark"
	"carizza/internal/pkg/apperror"
)

// UserRepository is a repository for the mark entity
type MarkRepository struct {
	repository
}

var _ mark.Repository = (*MarkRepository)(nil)

// New creates a new MarkRepository
func NewMarkRepository(repository *repository) (*MarkRepository, error) {
	return &MarkRepository{repository: *repository}, nil
}

func (r MarkRepository) autoMigrate() {
	if r.db.IsAutoMigrate() {
		r.db.DB().AutoMigrate(&mark.Mark{})
	}
}

// Get reads the album with the specified ID from the database.
func (r MarkRepository) Get(ctx context.Context, id uint) (*mark.Mark, error) {
	entity := &mark.Mark{}

	err := r.dbWithDefaults().First(entity, id).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return entity, apperror.ErrNotFound
		}
	}
	return entity, err
}

func (r MarkRepository) First(ctx context.Context, entity *mark.Mark) (*mark.Mark, error) {
	err := r.dbWithDefaults().Where(entity).First(entity).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return entity, apperror.ErrNotFound
		}
	}
	return entity, err
}

// Query retrieves the album records with the specified offset and limit from the database.
func (r MarkRepository) Query(ctx context.Context, cond selection_condition.SelectionCondition) ([]mark.Mark, error) {
	items := []mark.Mark{}
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
