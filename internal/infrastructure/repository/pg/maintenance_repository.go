package pg

import (
	"context"

	"github.com/jinzhu/gorm"

	"carizza/internal/pkg/apperror"

	"carizza/internal/domain"
	"carizza/internal/domain/maintenance"
)

// MaintenanceRepository is a repository for the service entity
type MaintenanceRepository struct {
	repository
}

var _ maintenance.Repository = (*MaintenanceRepository)(nil)

// NewMaintenanceRepository creates a new c
func NewMaintenanceRepository(repository *repository) (*MaintenanceRepository, error) {
	r := &MaintenanceRepository{repository: *repository}
	r.autoMigrate()
	return r, nil
}

func (r MaintenanceRepository) autoMigrate() {
	if r.db.IsAutoMigrate() {
		r.db.DB().AutoMigrate(
			&maintenance.Maintenance{},
			//&maintenance2work.Maintenance2Work{},
		)
	}
}

// Get reads the album with the specified ID from the database.
func (r MaintenanceRepository) Get(ctx context.Context, id uint) (*maintenance.Maintenance, error) {
	entity := &maintenance.Maintenance{}

	err := r.dbWithDefaults().First(entity, id).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return entity, apperror.ErrNotFound
		}
	}
	return entity, err
}

func (r MaintenanceRepository) First(ctx context.Context, entity *maintenance.Maintenance) (*maintenance.Maintenance, error) {
	err := r.dbWithDefaults().Where(entity).First(entity).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return entity, apperror.ErrNotFound
		}
	}
	return entity, err
}

// Query retrieves the album records with the specified offset and limit from the database.
func (r MaintenanceRepository) Query(ctx context.Context, cond domain.DBQueryConditions) ([]maintenance.Maintenance, error) {
	items := []maintenance.Maintenance{}
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
