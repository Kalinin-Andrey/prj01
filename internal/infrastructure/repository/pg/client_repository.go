package pg

import (
	"context"

	"github.com/jinzhu/gorm"

	"carizza/internal/pkg/apperror"
	minipkg_gorm "carizza/pkg/db/gorm"
	"carizza/pkg/selection_condition"

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

	err := r.DB().First(entity, id).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return entity, apperror.ErrNotFound
		}
	}
	return entity, err
}

func (r ClientRepository) First(ctx context.Context, entity *client.Client) (*client.Client, error) {
	err := r.DB().Where(entity).First(entity).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return entity, apperror.ErrNotFound
		}
	}
	return entity, err
}

// Query retrieves the album records with the specified offset and limit from the database.
func (r ClientRepository) Query(ctx context.Context, cond *selection_condition.SelectionCondition) ([]client.Client, error) {
	items := []client.Client{}
	db := minipkg_gorm.Conditions(r.DB(), cond)
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
