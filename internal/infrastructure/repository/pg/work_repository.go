package pg

import (
	"context"

	"github.com/jinzhu/gorm"

	minipkg_gorm "carizza/pkg/db/gorm"
	"carizza/pkg/selection_condition"

	"carizza/internal/pkg/apperror"

	"carizza/internal/domain/work"
)

// WorkRepository is a repository for the service entity
type WorkRepository struct {
	repository
}

var _ work.Repository = (*WorkRepository)(nil)

// NewWorkRepository creates a new c
func NewWorkRepository(repository *repository) (*WorkRepository, error) {
	r := &WorkRepository{repository: *repository}
	r.autoMigrate()
	return r, nil
}

func (r WorkRepository) autoMigrate() {
	if r.db.IsAutoMigrate() {
		r.db.DB().AutoMigrate(
			&work.Work{},
			//&work2supply.Work2Supply{},
		)
	}
}

// Get reads the album with the specified ID from the database.
func (r WorkRepository) Get(ctx context.Context, id uint) (*work.Work, error) {
	entity := &work.Work{}

	err := r.dbWithDefaults().First(entity, id).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return entity, apperror.ErrNotFound
		}
	}
	return entity, err
}

func (r WorkRepository) First(ctx context.Context, entity *work.Work) (*work.Work, error) {
	err := r.dbWithDefaults().Where(entity).First(entity).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return entity, apperror.ErrNotFound
		}
	}
	return entity, err
}

// Query retrieves the album records with the specified offset and limit from the database.
func (r WorkRepository) Query(ctx context.Context, cond selection_condition.SelectionCondition) ([]work.Work, error) {
	items := []work.Work{}
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
