package pg

import (
	"context"

	"github.com/jinzhu/gorm"

	"carizza/internal/pkg/apperror"

	"carizza/internal/domain"
	"carizza/internal/domain/order"
)

// OrderRepository is a repository for the service entity
type OrderRepository struct {
	repository
}

var _ order.Repository = (*OrderRepository)(nil)

// NewOrderRepository creates a new c
func NewOrderRepository(repository *repository) (*OrderRepository, error) {
	r := &OrderRepository{repository: *repository}
	r.autoMigrate()
	return r, nil
}

func (r OrderRepository) autoMigrate() {
	if r.db.IsAutoMigrate() {
		r.db.DB().AutoMigrate(&order.Order{})
	}
}

// Get reads the album with the specified ID from the database.
func (r OrderRepository) Get(ctx context.Context, id uint) (*order.Order, error) {
	entity := &order.Order{}

	err := r.dbWithDefaults().First(entity, id).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return entity, apperror.ErrNotFound
		}
	}
	r.setupEntityPeriod(entity)
	return entity, err
}

func (r OrderRepository) First(ctx context.Context, entity *order.Order) (*order.Order, error) {
	err := r.dbWithDefaults().Where(entity).First(entity).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return entity, apperror.ErrNotFound
		}
	}
	r.setupEntityPeriod(entity)
	return entity, err
}

// Query retrieves the album records with the specified offset and limit from the database.
func (r OrderRepository) Query(ctx context.Context, cond domain.DBQueryConditions) ([]order.Order, error) {
	items := []order.Order{}
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
	r.setupEntityPeriods(&items)
	return items, err
}

func (r OrderRepository) setupEntityPeriods(items *[]order.Order) {
	n := make([]order.Order, 0, len(*items))
	for _, item := range *items {
		r.setupEntityPeriod(&item)
		n = append(n, item)
	}
	*items = n
}

func (r OrderRepository) setupEntityPeriod(item *order.Order) {
	(*item).Period = order.Periods[item.PeriodID]
}
