package pg

import (
	"context"

	"github.com/jinzhu/gorm"

	minipkg_gorm "carizza/pkg/db/gorm"
	"carizza/pkg/selection_condition"

	"carizza/internal/pkg/apperror"

	"carizza/internal/domain/serie"
)

// UserRepository is a repository for the serie entity
type SerieRepository struct {
	repository
}

var _ serie.Repository = (*SerieRepository)(nil)

// New creates a new SerieRepository
func NewSerieRepository(repository *repository) (*SerieRepository, error) {
	return &SerieRepository{repository: *repository}, nil
}

func (r SerieRepository) autoMigrate() {
	if r.db.IsAutoMigrate() {
		r.db.DB().AutoMigrate(&serie.Serie{})
	}
}

// Get reads the album with the specified ID from the database.
func (r SerieRepository) Get(ctx context.Context, id uint) (*serie.Serie, error) {
	entity := &serie.Serie{}

	err := r.dbWithDefaults().First(entity, id).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return entity, apperror.ErrNotFound
		}
	}
	return entity, err
}

func (r SerieRepository) First(ctx context.Context, entity *serie.Serie) (*serie.Serie, error) {
	err := r.dbWithDefaults().Where(entity).First(entity).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return entity, apperror.ErrNotFound
		}
	}
	return entity, err
}

// Query retrieves the album records with the specified offset and limit from the database.
func (r SerieRepository) Query(ctx context.Context, cond selection_condition.SelectionCondition) ([]serie.Serie, error) {
	items := []serie.Serie{}
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
