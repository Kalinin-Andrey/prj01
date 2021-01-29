package pg

import (
	"context"

	"github.com/jinzhu/gorm"

	"carizza/internal/pkg/apperror"

	"carizza/internal/domain"
	"carizza/internal/domain/generation"
)

// UserRepository is a repository for the generation entity
type GenerationRepository struct {
	repository
}

var _ generation.Repository = (*GenerationRepository)(nil)

// NewGenerationRepository creates a new GenerationRepository
func NewGenerationRepository(repository *repository) (*GenerationRepository, error) {
	return &GenerationRepository{repository: *repository}, nil
}

func (r GenerationRepository) autoMigrate() {
	if r.db.IsAutoMigrate() {
		r.db.DB().AutoMigrate(&generation.Generation{})
	}
}

// Get reads the album with the specified ID from the database.
func (r GenerationRepository) Get(ctx context.Context, id uint) (*generation.Generation, error) {
	entity := &generation.Generation{}

	err := r.dbWithDefaults().First(entity, id).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return entity, apperror.ErrNotFound
		}
	}
	return entity, err
}

func (r GenerationRepository) First(ctx context.Context, entity *generation.Generation) (*generation.Generation, error) {
	err := r.dbWithDefaults().Where(entity).First(entity).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return entity, apperror.ErrNotFound
		}
	}
	return entity, err
}

// Query retrieves the album records with the specified offset and limit from the database.
func (r GenerationRepository) Query(ctx context.Context, cond domain.DBQueryConditions) ([]generation.Generation, error) {
	items := []generation.Generation{}
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
