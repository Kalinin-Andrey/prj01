package pg

import (
	"context"

	"github.com/jinzhu/gorm"

	"carizza/internal/pkg/apperror"

	"carizza/internal/domain"
	"carizza/internal/domain/client"
)

// ClientRepository is a repository for the service entity
type ClientRepository struct {
	repository
}

var _ client.Repository = (*ClientRepository)(nil)

// NewClientRepository creates a new c
func NewClientRepository(repository *repository) (*ClientRepository, error) {
	r := &ClientRepository{repository: *repository}
	r.autoMigrate()
	return r, nil
}

func (r ClientRepository) autoMigrate() {
	if r.db.IsAutoMigrate() {
		r.db.DB().AutoMigrate(&client.Client{})
	}
}

// Get reads the album with the specified ID from the database.
func (r ClientRepository) Get(ctx context.Context, id uint) (*client.Client, error) {
	entity := &client.Client{}

	err := r.dbWithDefaults().First(entity, id).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return entity, apperror.ErrNotFound
		}
	}
	return entity, err
}

func (r ClientRepository) First(ctx context.Context, entity *client.Client) (*client.Client, error) {
	err := r.dbWithDefaults().Where(entity).First(entity).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return entity, apperror.ErrNotFound
		}
	}
	return entity, err
}

// Query retrieves the album records with the specified offset and limit from the database.
func (r ClientRepository) Query(ctx context.Context, cond domain.DBQueryConditions) ([]client.Client, error) {
	items := []client.Client{}
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
